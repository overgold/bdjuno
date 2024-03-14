package stake

import (
	"git.ooo.ua/vipcoin/ovg-chain/x/stake/types"
	juno "github.com/forbole/juno/v5/types"
)

// handleMsgCreateSystemStakeAccountAddress allows to properly handle a message
func (m *Module) handleMsgManageSystemStake(tx *juno.Tx, _ int, msg *types.MsgManageSystemStake) error {
	return m.stakeRepo.InsertMsgManageSystemStake(tx.TxHash, *msg)
}
