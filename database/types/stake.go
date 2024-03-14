package types

type (
	// StakeMsgSell - db model for 'overgold_stake_sell'
	StakeMsgSell struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgSellCancel - db model for 'overgold_stake_sell_cancel'
	StakeMsgSellCancel struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgBuy - db model for 'overgold_stake_buy'
	StakeMsgBuy struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgDistribute - db model for 'overgold_stake_distribute_rewards'
	StakeMsgDistribute struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
	}

	// StakeMsgClaim - db model for 'overgold_stake_claim_reward'
	StakeMsgClaim struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgTransferFromUser - db model for 'overgold_stake_transfer_from_user'
	StakeMsgTransferFromUser struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
		Address string `db:"address"`
	}

	// StakeMsgTransferToUser - db model for 'overgold_stake_transfer_to_user'
	StakeMsgTransferToUser struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
		Address string `db:"address"`
	}

	// StakeMsgCreateSystemStakeAccountAddress - db model for 'overgold_stake_create_system_stake_account_address'
	StakeMsgCreateSystemStakeAccountAddress struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Address string `db:"address"`
	}

	// StakeMsgUpdateSystemStakeAccountAddress - db model for 'overgold_stake_update_system_stake_account_address'
	StakeMsgUpdateSystemStakeAccountAddress struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Address string `db:"address"`
	}

	// StakeMsgDeleteSystemStakeAccountAddress - db model for 'overgold_stake_delete_system_stake_account_address'
	StakeMsgDeleteSystemStakeAccountAddress struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
	}

	// StakeMsgManageSystemStake - db model for 'overgold_stake_manage_system_stake'
	StakeMsgManageSystemStake struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
		Kind    string `db:"kind"`
	}
)
