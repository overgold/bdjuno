package stake

import (
	"git.ooo.ua/vipcoin/ovg-chain/x/stake/types"
	juno "github.com/forbole/juno/v5/types"
)

// handleMsgCreateSystemStakeAccountAddress allows to properly handle a message
func (m *Module) handleMsgCreateSystemStakeAccountAddress(tx *juno.Tx, _ int, msg *types.MsgCreateSystemStakeAccountAddress) error {
	return m.stakeRepo.InsertMsgCreateSystemStakeAccountAddress(tx.TxHash, *msg)
}

// handleMsgUpdateSystemStakeAccountAddress allows to properly handle a message
func (m *Module) handleMsgUpdateSystemStakeAccountAddress(tx *juno.Tx, _ int, msg *types.MsgUpdateSystemStakeAccountAddress) error {
	return m.stakeRepo.InsertMsgUpdateSystemStakeAccountAddress(tx.TxHash, *msg)
}

// handleMsgDeleteSystemStakeAccountAddress allows to properly handle a message
func (m *Module) handleMsgDeleteSystemStakeAccountAddress(tx *juno.Tx, _ int, msg *types.MsgDeleteSystemStakeAccountAddress) error {
	return m.stakeRepo.InsertMsgDeleteSystemStakeAccountAddress(tx.TxHash, *msg)
}
