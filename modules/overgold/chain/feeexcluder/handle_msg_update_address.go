package feeexcluder

import (
	"git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	juno "github.com/forbole/juno/v5/types"
)

// handleMsgUpdateAddress allows to properly handle a message
func (m *Module) handleMsgUpdateAddress(tx *juno.Tx, index int, msg *types.MsgUpdateAddress) error {
	// TODO
	return nil
}
