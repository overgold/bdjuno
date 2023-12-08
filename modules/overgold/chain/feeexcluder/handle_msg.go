package feeexcluder

import (
	"git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch feeExcluderMsg := msg.(type) {
	case *types.MsgCreateAddress:
		return m.handleMsgCreateAddress(tx, index, feeExcluderMsg)
	case *types.MsgUpdateAddress:
		return m.handleMsgUpdateAddress(tx, index, feeExcluderMsg)
	case *types.MsgDeleteAddress:
		return m.handleMsgDeleteAddress(tx, index, feeExcluderMsg)
	case *types.MsgCreateTariffs:
		// TODO: add logic and other cases
		return nil
	case *types.MsgUpdateTariffs:
		// TODO: add logic and other cases
		return nil
	case *types.MsgDeleteTariffs:
		// TODO: add logic and other cases
		return nil
	default:
		return nil
	}
}
