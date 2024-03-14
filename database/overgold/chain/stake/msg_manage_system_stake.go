package stake

import (
	"git.ooo.ua/vipcoin/lib/errs"
	stake "git.ooo.ua/vipcoin/ovg-chain/x/stake/types"

	"github.com/forbole/bdjuno/v4/database/overgold/chain"
)

// InsertMsgManageSystemStake - insert a new MsgManageSystemStake in a database (overgold_stake_manage_system_stake).
func (r Repository) InsertMsgManageSystemStake(hash string, msgs ...stake.MsgManageSystemStake) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO overgold_stake_manage_system_stake (tx_hash, creator, amount, kind) 
		VALUES ( $1, $2, $3, $4 )
	`

	for _, msg := range msgs {
		m, err := toMsgManageSystemStakeDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator, m.Amount, m.Kind); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
