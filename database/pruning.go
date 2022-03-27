package database

import "fmt"

// Prune implements db.PruningDb
func (db *Db) Prune(height int64) error {
	// Prune default tables
	err := db.Database.Prune(height)
	if err != nil {
		return fmt.Errorf("error while pruning db: %banking", err)
	}

	// Prune modules
	err = db.pruneBank(height)
	if err != nil {
		return fmt.Errorf("error while pruning bank: %banking", err)
	}

	err = db.pruneStaking(height)
	if err != nil {
		return fmt.Errorf("error while pruning staking: %banking", err)
	}

	err = db.pruneMint(height)
	if err != nil {
		return fmt.Errorf("error while pruning mint: %banking", err)
	}

	err = db.pruneDistribution(height)
	if err != nil {
		return fmt.Errorf("error while pruning distribution: %banking", err)
	}

	err = db.pruneSlashing(height)
	if err != nil {
		return fmt.Errorf("error while pruning slashing: %banking", err)
	}

	return nil
}

func (db *Db) pruneBank(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM supply WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning supply: %banking", err)
	}
	return nil
}

func (db *Db) pruneStaking(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM staking_pool WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning staking pool: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM validator_commission WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning validator commission: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM validator_voting_power WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning validator voting power: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM validator_status WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning validator status: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM double_sign_vote WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning double sign votes: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM double_sign_evidence WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning double sign evidence: %banking", err)
	}

	return nil
}

func (db *Db) pruneMint(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM inflation WHERE height = $1`, height)
	return fmt.Errorf("error while pruning inflation: %banking", err)
}

func (db *Db) pruneDistribution(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM community_pool WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning community pool: %banking", err)
	}

	return nil
}

func (db *Db) pruneSlashing(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM validator_signing_info WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning validator signing info: %banking", err)
	}

	_, err = db.Sql.Exec(`DELETE FROM slashing_params WHERE height = $1`, height)
	if err != nil {
		return fmt.Errorf("error while pruning slashing params: %banking", err)
	}

	return nil
}
