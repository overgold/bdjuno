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

// GetAllMsgCreateTariffs - method that get data from a db (overgold_feeexcluder_create_tariffs).
// TODO: use JOIN and other db model
func (r Repository) GetAllMsgCreateTariffs(f filter.Filter) ([]fe.MsgCreateTariffs, error) {
	q, args := f.Build(tableCreateTariffs)

	// 1) get create tariffs
	var createTariffs []types.FeeExcluderCreateTariffs
	if err := r.db.Select(&createTariffs, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableCreateTariffs}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(createTariffs) == 0 {
		return nil, errs.NotFound{What: tableCreateTariffs}
	}

	result := make([]fe.MsgCreateTariffs, 0, len(createTariffs))
	for _, ct := range createTariffs {
		// 2) get tariff
		tariff, err := r.GetAllTariff(filter.NewFilter().SetArgument(types.FieldID, ct.TariffID))
		if err != nil {
			return nil, err
		}
		if len(tariff) == 0 {
			return nil, errs.NotFound{What: tableTariff}
		}

		result = append(result, toMsgCreateTariffsDomain(ct, tariff[0]))
	}

	return result, nil
}

// InsertToMsgCreateTariffs - insert new data in a database (overgold_feeexcluder_create_tariffs).
func (r Repository) InsertToMsgCreateTariffs(hash string, ct fe.MsgCreateTariffs) error {
	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// 1) add tariff
	tariffID, err := r.InsertToTariff(tx, ct.Tariff)
	if err != nil {
		return err
	}

	// 2) add create tariffs
	q := `
		INSERT INTO overgold_feeexcluder_create_tariffs (
			tx_hash, creator, denom, tariff_id
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, tx_hash, creator, denom, tariff_id
	`

	m := toMsgCreateTariffsDatabase(hash, 0, tariffID, ct)
	if _, err = tx.Exec(q, m.TxHash, m.Creator, m.Denom, m.TariffID); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	if err = tx.Commit(); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}

// / TODO
// UpdateMsgCreateTariffs - method that updates in a database (overgold_feeexcluder_create_tariffs).
func (r Repository) UpdateMsgCreateTariffs(hash string, id uint64, ct fe.MsgCreateTariffs) error {
	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer commit(tx, err)

	// 1) get unique tariff id
	tariffList, err := r.getAllTariffWithUniqueID(filter.NewFilter().SetArgument(types.FieldMsgID, ct.Tariff.Id))
	if err != nil {
		return err
	}
	if len(tariffList) != 1 {
		return errs.Internal{Cause: "expected 1 tariff"}
	}
	tariffID := tariffList[0].ID

	// 2) update create tariffs
	q := `UPDATE overgold_feeexcluder_create_tariffs SET
				 tx_hash = $1,
				 creator = $2,
				 denom = $3,
            	 tariff_id = $4                  
			 WHERE id = $5`

	m := toMsgCreateTariffsDatabase(hash, id, tariffID, ct)
	if _, err = tx.Exec(q, m.TxHash, m.Creator, m.Denom, m.TariffID, m.ID); err != nil {
		return err
	}

	// 3) get unique id for tariff
	tariffs, err := r.GetAllTariffs(filter.NewFilter().SetArgument(types.FieldID, ct.Tariff.Id))
	if err != nil {
		return err
	}

	if len(tariffs) != 1 {
		return errs.Internal{Cause: "expected 1 tariff"}
	}
	tariff := tariffs[0].Tariffs

	// 4) update tariff
	// TODO: dont't use it (testing)
	if err = r.UpdateTariff(tx, tariffID, tariff[0]); err != nil {
		return err
	}

	return nil
}

// DeleteMsgCreateTariffs - method that deletes data in a database (overgold_feeexcluder_create_tariffs).
func (r Repository) DeleteMsgCreateTariffs(id uint64) error {
	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer commit(tx, err)

	// 1) delete create tariffs
	// 1.a) get ids for tariff
	messages, err := r.GetAllMsgCreateTariffs(filter.NewFilter().SetArgument(types.FieldID, id))
	if err != nil {
		if !errors.As(err, &errs.NotFound{}) {
			return err
		}
	}

	// 1.b) exec delete create tariffs
	q := `DELETE FROM overgold_feeexcluder_create_tariffs WHERE id IN ($1)`

	if _, err = tx.Exec(q, id); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	// 2) delete tariff
	for _, msg := range messages {
		if err = r.DeleteTariff(tx, msg.Tariff.Id); err != nil {
			return err
		}
	}

	return nil
}
