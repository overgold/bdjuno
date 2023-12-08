package feeexcluder

import (
	"context"
	"database/sql"
	"errors"

	"git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"
	fe "git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"

	"github.com/forbole/bdjuno/v4/database/types"
)

// GetAllMsgUpdateTariffs - method that get data from a db (overgold_feeexcluder_update_tariffs).
// TODO: use JOIN and other db model
func (r Repository) GetAllMsgUpdateTariffs(f filter.Filter) ([]fe.MsgUpdateTariffs, error) {
	q, args := f.Build(tableUpdateTariffs)

	// 1) get update tariffs
	var updateTariffs []types.FeeExcluderUpdateTariffs
	if err := r.db.Select(&updateTariffs, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableUpdateTariffs}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(updateTariffs) == 0 {
		return nil, errs.NotFound{What: tableUpdateTariffs}
	}

	// 2) get tariff
	result := make([]fe.MsgUpdateTariffs, 0, len(updateTariffs))
	for _, ut := range updateTariffs {
		tariff, err := r.GetAllTariff(filter.NewFilter().SetArgument(types.FieldID, ut.TariffID))
		if err != nil {
			return nil, err
		}
		if len(tariff) == 0 {
			return nil, errs.NotFound{What: tableTariff}
		}

		result = append(result, toMsgUpdateTariffsDomain(ut, tariff[0]))
	}

	return result, nil
}

// InsertToMsgUpdateTariffs - insert new data in a database (overgold_feeexcluder_update_tariffs).
func (r Repository) InsertToMsgUpdateTariffs(hash string, ut ...fe.MsgUpdateTariffs) error {
	if len(ut) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// 1) add tariff
	for _, t := range ut {
		if _, err = r.InsertToTariff(tx, t.Tariff); err != nil {
			return err
		}
	}

	// 2) add update tariffs

	q := `
		INSERT INTO overgold_feeexcluder_update_tariffs (
			tx_hash, creator, denom, tariff_id
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, tx_hash, creator, denom, tariff_id
	`

	for _, t := range ut {
		m := toMsgUpdateTariffsDatabase(hash, 0, t)
		if _, err = tx.Exec(q, m.TxHash, m.Creator, m.Denom, m.TariffID); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	if err = tx.Commit(); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}

// UpdateMsgUpdateTariffs - method that updates in a database (overgold_feeexcluder_update_tariffs).
func (r Repository) UpdateMsgUpdateTariffs(hash string, id uint64, ut ...fe.MsgUpdateTariffs) error {
	if len(ut) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// 1) update update tariffs
	q := `UPDATE overgold_feeexcluder_update_tariffs SET
				 tx_hash = $1,
				 creator = $2,
            	 tariff_id = $3,
            	 denom = $4
			 WHERE id = $5`

	for _, t := range ut {
		m := toMsgUpdateTariffsDatabase(hash, id, t)
		if _, err = tx.Exec(q, m.TxHash, m.Creator, m.TariffID, m.Denom, m.ID); err != nil {
			return err
		}
	}

	// 2) update tariff
	for _, t := range ut {
		if err = r.UpdateTariff(tx, t.Tariff); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteMsgUpdateTariffs - method that deletes data in a database (overgold_feeexcluder_update_tariffs).
func (r Repository) DeleteMsgUpdateTariffs(id uint64) error {
	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer tx.Rollback()

	// 1) delete update tariffs
	// 1.a) get ids for tariff
	messages, err := r.GetAllMsgUpdateTariffs(filter.NewFilter().SetArgument(types.FieldID, id))
	if err != nil {
		if !errors.As(err, &errs.NotFound{}) {
			return err
		}
	}

	// 1.b) exec delete update tariffs
	q := `DELETE FROM overgold_feeexcluder_update_tariffs WHERE id IN ($1)`

	if _, err = tx.Exec(q, id); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	// 2) delete tariff
	for _, msg := range messages {
		if err = r.DeleteTariff(tx, msg.Tariff.Id); err != nil {
			return err
		}
	}

	return tx.Commit()
}
