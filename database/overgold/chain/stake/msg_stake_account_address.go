package stake

import (
	"git.ooo.ua/vipcoin/lib/errs"
	stake "git.ooo.ua/vipcoin/ovg-chain/x/stake/types"

	"github.com/forbole/bdjuno/v4/database/overgold/chain"
)

// InsertMsgCreateSystemStakeAccountAddress - insert a new MsgCreateSystemStakeAccountAddress
// in a database (overgold_stake_create_system_stake_account_address).
func (r Repository) InsertMsgCreateSystemStakeAccountAddress(hash string, msgs ...stake.MsgCreateSystemStakeAccountAddress) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO overgold_stake_create_system_stake_account_address (tx_hash, creator, address) 
		VALUES ($1, $2, $3)
	`

	for _, msg := range msgs {
		m, err := toMsgCreateSystemStakeAccountAddressDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator, m.Address); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// InsertMsgUpdateSystemStakeAccountAddress - insert a new MsgUpdateSystemStakeAccountAddress
// in a database (overgold_stake_update_system_stake_account_address).
func (r Repository) InsertMsgUpdateSystemStakeAccountAddress(hash string, msgs ...stake.MsgUpdateSystemStakeAccountAddress) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO overgold_stake_update_system_stake_account_address (tx_hash, creator, address) 
		VALUES ($1, $2, $3)
	`

	for _, msg := range msgs {
		m, err := toMsgUpdateSystemStakeAccountAddressDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator, m.Address); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// InsertMsgDeleteSystemStakeAccountAddress - insert a new MsgDeleteSystemStakeAccountAddress
// in a database (overgold_stake_delete_system_stake_account_address).
func (r Repository) InsertMsgDeleteSystemStakeAccountAddress(hash string, msgs ...stake.MsgDeleteSystemStakeAccountAddress) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO overgold_stake_delete_system_stake_account_address (tx_hash, creator) 
		VALUES ($1, $2)
	`

	for _, msg := range msgs {
		m, err := toMsgDeleteSystemStakeAccountAddressDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
