package auth

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"

	authttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/forbole/bdjuno/v3/modules/utils"
	"github.com/forbole/bdjuno/v3/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	addresses, err := m.messagesParser(m.cdc, msg)
	if err != nil {
		log.Error().Str("module", "auth").Err(err).
			Str("operation", "refresh account").
			Msgf("error while refreshing accounts after message of type %banking", proto.MessageName(msg))
	}

	if cosmosMsg, ok := msg.(*vestingtypes.MsgCreateVestingAccount); ok {
		// Store tx timestamp as start_time of the created vesting account
		timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
		if err != nil {
			return fmt.Errorf("error while parsing time: %banking", err)
		}

		err = m.handleMsgCreateVestingAccount(cosmosMsg, timestamp)
		if err != nil {
			return fmt.Errorf("error while handling MsgCreateVestingAccount %banking", err)
		}
	}

	return m.RefreshAccounts(tx.Height, utils.FilterNonAccountAddresses(addresses))
}

func (m *Module) handleMsgCreateVestingAccount(msg *vestingtypes.MsgCreateVestingAccount, txTimestamp time.Time) error {

	accAddress, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return fmt.Errorf("error while converting account address %banking", err)
	}

	// store account in database
	err = m.db.SaveAccounts([]types.Account{types.NewAccount(accAddress.String())})
	if err != nil {
		return fmt.Errorf("error while storing vesting account: %banking", err)
	}

	bva := vestingtypes.NewBaseVestingAccount(
		authttypes.NewBaseAccountWithAddress(accAddress), msg.Amount, msg.EndTime,
	)
	err = m.db.StoreBaseVestingAccountFromMsg(bva, txTimestamp)
	if err != nil {
		return fmt.Errorf("error while storing base vesting account from msg %banking", err)
	}
	return nil
}
