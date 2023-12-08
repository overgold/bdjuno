package feeexcluder

import (
	"context"
	"database/sql"
	"errors"

	"git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"
	fe "git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	"github.com/jmoiron/sqlx"

	"github.com/forbole/bdjuno/v4/database/types"
)

// GetAllTariff - method that get data from a db (overgold_feeexcluder_tariff).
// TODO: use JOIN and other db model, e.g.:
//
// SELECT f.*, mt.*
// FROM overgold_feeexcluder_fees AS f
// JOIN overgold_feeexcluder_m2m_tariff_fees AS mt ON f.id = mt.fees_id
// JOIN overgold_feeexcluder_tariff AS t ON mt.tariff_id = t.id;
func (r Repository) GetAllTariff(f filter.Filter) ([]*fe.Tariff, error) {
	q, args := f.Build(tableTariff)

	// 1) get tariff
	var tariffs []types.FeeExcluderTariff
	if err := r.db.Select(&tariffs, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableTariff}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(tariffs) == 0 {
		return nil, errs.NotFound{What: tableTariff}
	}

	// TODO: refactor, use JOIN and other db model, e.g: overgold_feeexcluder_m2m_tariff_fees.fees_id...
	result := make([]*fe.Tariff, 0, len(tariffs))
	for _, t := range tariffs {
		// 2) get m2m tariff fees
		m2mFees, err := r.GetAllM2MTariffFees(filter.NewFilter().SetArgument(types.FieldTariffID, t.ID))
		if err != nil {
			return nil, err
		}

		feeIDs := make([]uint64, 0, len(m2mFees))
		for _, m2m := range m2mFees {
			feeIDs = append(feeIDs, m2m.FeesID)
		}

		// 3) get fees
		fees, err := r.GetAllFees(filter.NewFilter().SetArgument(types.FieldID, feeIDs))
		if err != nil {
			return nil, err
		}

		result = append(result, toTariffDomain(t, fees))
	}

	return result, nil
}

// InsertToTariff - insert new data in a database (overgold_feeexcluder_tariff).
func (r Repository) InsertToTariff(tx *sqlx.Tx, tariff *fe.Tariff) (lastID uint64, err error) {
	if tx == nil {
		tx, err = r.db.BeginTxx(context.Background(), &sql.TxOptions{})
		if err != nil {
			return 0, errs.Internal{Cause: err.Error()}
		}

		defer commit(tx, err)
	}

	// 1) add tariff
	q := `
		INSERT INTO overgold_feeexcluder_tariff (
			id, amount, denom, min_ref_balance
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id
	`

	m, err := toTariffDatabase(tariff)
	if err != nil {
		return 0, errs.Internal{Cause: err.Error()}
	}

	if err = tx.QueryRowx(q, m.ID, m.Amount, m.Denom, m.MinRefBalance).Scan(&lastID); err != nil {
		return 0, errs.Internal{Cause: err.Error()}
	}

	// 2) add fees
	for _, f := range tariff.Fees {
		if _, err = r.InsertToFees(tx, f); err != nil {
			return 0, err
		}
	}

	// 3) add many-to-many tariff fees
	m2m := make([]types.FeeExcluderM2MTariffFees, 0, len(tariff.Fees))
	for _, f := range tariff.Fees {
		m2m = append(m2m, types.FeeExcluderM2MTariffFees{
			TariffID: tariff.Id,
			FeesID:   f.Id,
		})
	}

	return lastID, r.InsertToM2MTariffFees(tx, m2m...)
}

// UpdateTariff - method that updates in a database (overgold_feeexcluder_tariff).
func (r Repository) UpdateTariff(tx *sqlx.Tx, tariffList ...*fe.Tariff) (err error) {
	if len(tariffList) == 0 {
		return nil
	}

	if tx == nil {
		tx, err = r.db.BeginTxx(context.Background(), &sql.TxOptions{})
		if err != nil {
			return errs.Internal{Cause: err.Error()}
		}

		defer commit(tx, err)
	}

	// 1) update tariff
	q := `UPDATE overgold_feeexcluder_tariff SET
				 amount = $1,
				 denom = $2,
				 min_ref_balance = $3
			 WHERE id = $4`

	for _, t := range tariffList {
		m, err := toTariffDatabase(t)
		if err != nil {
			return errs.Internal{Cause: err.Error()}
		}

		if _, err = tx.Exec(q, m.Amount, m.Denom, m.MinRefBalance, m.ID); err != nil {
			return err
		}
	}

	// 2) update fees
	for _, t := range tariffList {
		if err = r.UpdateFees(tx, t.Fees...); err != nil {
			return err
		}
	}

	return nil
}

// DeleteTariff - method that deletes data in a database (overgold_feeexcluder_tariff).
func (r Repository) DeleteTariff(tx *sqlx.Tx, id uint64) (err error) {
	if tx == nil {
		tx, err = r.db.BeginTxx(context.Background(), &sql.TxOptions{})
		if err != nil {
			return errs.Internal{Cause: err.Error()}
		}

		defer commit(tx, err)
	}

	// 1) delete many-to-many tariff fees and get ids
	m2m, err := r.GetAllM2MTariffFees(filter.NewFilter().SetArgument(types.FieldTariffID, id))
	if err != nil {
		if !errors.As(err, &errs.NotFound{}) {
			return err
		}
	}

	if err = r.DeleteM2MTariffFeesByTariff(tx, id); err != nil {
		return err
	}

	// 2) delete fees
	for _, m := range m2m {
		if err = r.DeleteFees(tx, m.FeesID); err != nil {
			return err
		}
	}

	// 3) delete tariff
	q := `DELETE FROM overgold_feeexcluder_tariff WHERE id IN ($1)`

	if _, err = tx.Exec(q, id); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}
