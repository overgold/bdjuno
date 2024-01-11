package types

import (
	"database/sql"
	"time"
)

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
)

// BLOCK GENESIS STATE
type (
	// StakeGenesisState - db model for 'overgold_stake_genesis_state'
	StakeGenesisState struct {
		ID              uint64 `db:"id"` // unique id as primary key
		DailyStatsCount uint64 `db:"daily_stats_count"`
		TotalFree       uint64 `db:"total_free"`
		TotalSell       uint64 `db:"total_sell"`
		UsersCount      uint64 `db:"users_count"`
	}

	// ~~~~~ L2 ~~~~~

	// StakeStats - db model for 'overgold_stake_stats'
	StakeStats struct {
		ID      uint64        `db:"id"`       // unique id as primary key
		StatsID sql.NullInt64 `db:"stats_id"` // id for *StakeDailyStats
		Date    time.Time     `db:"date"`
	}

	// StakeStakes - db model for 'overgold_stake_stakes'
	StakeStakes struct {
		ID             uint64        `db:"id"`       // unique id as primary key
		StakeID        sql.NullInt64 `db:"stake_id"` // id for *StakeStake
		AccountAddress string        `db:"account_address"`
	}

	// StakeBuyRequests - db model for 'overgold_stake_buy_requests'
	StakeBuyRequests struct {
		ID             uint64  `db:"id"` // unique id as primary key
		AccountAddress string  `db:"account_address"`
		Amount         DbCoins `db:"amount"` // used without separate table
	}

	// StakeSellRequests - db model for 'overgold_stake_sell_requests'
	StakeSellRequests struct {
		ID             uint64  `db:"id"` // unique id as primary key
		AccountAddress string  `db:"account_address"`
		Amount         DbCoins `db:"amount"` // used without separate table
	}

	// StakeLimits - db model for 'overgold_stake_stake_limits'
	StakeLimits struct {
		ID           uint64        `db:"id"`             // unique id as primary key
		MinAmountsID sql.NullInt64 `db:"min_amounts_id"` // id for *StakeMinAmounts
		Creator      string        `db:"creator"`
		Denom        string        `db:"denom"`
	}

	// StakeUserStats - db model for 'overgold_stake_user_stats'
	StakeUserStats struct {
		ID            uint64        `db:"id"`              // unique id as primary key
		UniqueUsersID sql.NullInt64 `db:"unique_users_id"` // id for *StakeUniqueUsers
		Date          time.Time     `db:"date"`
	}

	// StakeSellParams - db model for 'overgold_stake_sell_params'
	StakeSellParams struct {
		ID            uint64        `db:"id"`              // unique id as primary key
		StakeParamsID sql.NullInt64 `db:"stake_params_id"` // id for *StakeSellStakeParams
		Creator       string        `db:"creator"`
		Denom         string        `db:"denom"`
	}

	// ~~~~~ L3 ~~~~~

	// StakeDailyStats - db model for 'overgold_stake_daily_stats'
	StakeDailyStats struct {
		ID                  uint64  `db:"id"`     // unique id as primary key
		MsgID               uint64  `db:"msg_id"` // daily stats id from message
		CountBurn           uint64  `db:"count_burn"`
		CountBuy            uint64  `db:"count_buy"`
		CountIssue          uint64  `db:"count_issue"`
		CountSell           uint64  `db:"count_sell"`
		CountWithdraw       uint64  `db:"count_withdraw"`
		CountUsers          uint64  `db:"count_users"`
		CountUsersBurn      uint64  `db:"count_users_burn"`
		CountUsersBuy       uint64  `db:"count_users_buy"`
		CountUsersIssue     uint64  `db:"count_users_issue"`
		CountUsersSell      uint64  `db:"count_users_sell"`
		CountUsersWithdraw  uint64  `db:"count_users_withdraw"`
		DistributionsByUser uint64  `db:"distributions_by_user"`
		AmountBurn          DbCoins `db:"amount_burn"`
		AmountBuy           DbCoins `db:"amount_buy"`
		AmountIssue         DbCoins `db:"amount_issue"`
		AmountSell          DbCoins `db:"amount_sell"`
		AmountWithdraw      DbCoins `db:"amount_withdraw"`
		DistributionsTotal  DbCoins `db:"distributions_total"`
	}

	// StakeStake - db model for 'overgold_stake_stake'
	StakeStake struct {
		ID           uint64  `db:"id"` // unique id as primary key
		RewardAmount DbCoins `db:"reward_amount"`
		SellAmount   DbCoins `db:"sell_amount"`
		StakeAmount  DbCoins `db:"stake_amount"`
	}

	// StakeMinAmounts - db model for 'overgold_stake_min_amounts'
	StakeMinAmounts struct {
		ID   uint64  `db:"id"` // unique id as primary key
		Sell DbCoins `db:"sell"`
		Buy  DbCoins `db:"buy"`
	}

	// StakeUniqueUsers - db model for 'overgold_stake_unique_users'
	StakeUniqueUsers struct {
		ID       uint64   `db:"id"` // unique id as primary key
		Burn     []string `db:"burn"`
		Buy      []string `db:"buy"`
		Issue    []string `db:"issue"`
		Sell     []string `db:"sell"`
		Total    []string `db:"total"`
		Withdraw []string `db:"withdraw"`
	}

	// StakeSellStakeParams - db model for 'overgold_stake_sell_stake_params'
	StakeSellStakeParams struct {
		ID              uint64 `db:"id"` // unique id as primary key
		MinSellRequests uint64 `db:"min_sell_requests"`
		MaxSellRequests uint64 `db:"max_sell_requests"`
	}

	// ~~~~~ L4 ~~~~~

	// StakeTransaction - db model for 'overgold_stake_transaction'
	StakeTransaction struct {
		ID          uint64    `db:"id"` // unique id as primary key
		Kind        int32     `db:"kind"`
		AddressFrom string    `db:"address_from"`
		AddressTo   string    `db:"address_to"`
		Description string    `db:"description"`
		Hash        string    `db:"hash"`
		Timestamp   time.Time `db:"timestamp"`
		Amount      DbCoins   `db:"amount"`
	}

	// ~~~~~ M2M ~~~~~

	// StakeM2MGenesisStateStats - db model for 'overgold_stake_m2m_genesis_state_stats'
	StakeM2MGenesisStateStats struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		StatsID        uint64 `db:"stats_id"`
	}

	// StakeM2MGenesisStateStakes - db model for 'overgold_stake_m2m_genesis_state_stakes'
	StakeM2MGenesisStateStakes struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		StakesID       uint64 `db:"stakes_id"`
	}

	// StakeM2MGenesisStateBuyRequests - db model for 'overgold_stake_m2m_genesis_state_buy_requests'
	StakeM2MGenesisStateBuyRequests struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		BuyRequestsID  uint64 `db:"buy_requests_id"`
	}

	// StakeM2MGenesisStateSellRequests - db model for 'overgold_stake_m2m_genesis_state_sell_requests'
	StakeM2MGenesisStateSellRequests struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		SellRequestsID uint64 `db:"sell_requests_id"`
	}

	// StakeM2MGenesisStakeStateLimits - db model for 'overgold_stake_m2m_genesis_state_stake_limits'
	StakeM2MGenesisStakeStateLimits struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		StakeLimitsID  uint64 `db:"stake_limits_id"`
	}

	// StakeM2MGenesisStateUserStats - db model for 'overgold_stake_m2m_genesis_state_user_stats'
	StakeM2MGenesisStateUserStats struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		UserStatsID    uint64 `db:"user_stats_id"`
	}

	// StakeM2MGenesisStateSellParams - db model for 'overgold_stake_m2m_genesis_state_sell_params'
	StakeM2MGenesisStateSellParams struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		SellParamsID   uint64 `db:"sell_params_id"`
	}

	// StakeM2MStakeTransactions - db model for 'overgold_stake_m2m_stake_transaction'
	StakeM2MStakeTransactions struct {
		GenesisStateID uint64 `db:"genesis_state_id"`
		SellParamsID   uint64 `db:"sell_params_id"`
	}
)
