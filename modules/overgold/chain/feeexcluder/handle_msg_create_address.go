package feeexcluder

import (
	"git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	juno "github.com/forbole/juno/v5/types"
)

// handleMsgCreateAddress allows to properly handle a message
func (m *Module) handleMsgCreateAddress(tx *juno.Tx, index int, msg *types.MsgCreateAddress) error {
	// TODO
	return nil
}
