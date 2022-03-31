package accounts

import (
	"git.ooo.ua/vipcoin/chain/x/accounts/types"
	"git.ooo.ua/vipcoin/lib/filter"
	juno "github.com/forbole/juno/v2/types"
)

// handleMsgSetState allows to properly handle a handleMsgSetState
func (m *Module) handleMsgSetState(tx *juno.Tx, index int, msg *types.MsgSetState) error {
	if err := m.accountRepo.SaveState(msg); err != nil {
		return err
	}

	acc, err := m.accountRepo.GetAccounts(filter.NewFilter().SetArgument(FieldHash, msg.Hash))
	if err != nil {
		return err
	}

	if len(acc) != 1 {
		return types.ErrInvalidHashField
	}

	acc[0].State = msg.State

	return m.accountRepo.UpdateAccounts(acc...)
}
