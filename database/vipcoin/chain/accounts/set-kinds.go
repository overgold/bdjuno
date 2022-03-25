/*
 * Copyright 2022 Business Process Technologies. All rights reserved.
 */

package accounts

import (
	"context"
	"database/sql"

	accountstypes "git.ooo.ua/vipcoin/chain/x/accounts/types"
	"git.ooo.ua/vipcoin/lib/filter"
	"github.com/forbole/bdjuno/v2/database/types"
)

func (r Repository) SaveSetKinds(msg ...*accountstypes.MsgSetKinds) error {
	if len(msg) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO vipcoin_chain_accounts_set_kinds 
			(creator, hash, kinds) 
			VALUES 
			(:creator, :hash, :kinds)`

	for _, kinds := range msg {
		if _, err := tx.NamedExec(query, toSetKindsDatabase(kinds)); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r Repository) GetSetKinds(accfilter filter.Filter) ([]*accountstypes.MsgSetKinds, error) {
	query, args := accfilter.Build("vipcoin_chain_accounts_set_kinds",
		`creator, hash, kinds`)

	var result []types.DBSetKinds
	if err := r.db.Select(&result, query, args...); err != nil {
		return []*accountstypes.MsgSetKinds{}, err
	}

	kinds := make([]*accountstypes.MsgSetKinds, 0, len(result))
	for _, kind := range result {
		kinds = append(kinds, toSetKindsDomain(kind))
	}

	return kinds, nil
}
