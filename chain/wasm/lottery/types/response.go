package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AllCombinationResponse struct {
	Combination []string `json:"combination"`
}

type WinnerRewardClaims struct {
	Claimed bool   `json:"claimed"`
	Ranks   []uint `json:"ranks"`
}

type Winner struct {
	Address string             `json:"address"`
	Claims  WinnerRewardClaims `json:"claims"`
}

type AllWinnerResponse struct {
	Winners []Winner `json:"winners"`
}

type ConfigResponse struct {
	Admin                string   `json:"admin"`
	BlockTimePlay        uint64   `json:"block_time_play"`
	CombinationLen       uint8    `json:"combination_len"`
	EveryBlockTimePlay   uint64   `json:"every_block_time_play"`
	JackpotPercent       uint8    `json:"jackpot_percent"`
	LotteryCounter       uint64   `json:"lottery_counter"`
	LotteryCw20Addr      string   `json:"lottery_cw20_addr"`
	LotteryWorkerFee     uint8    `json:"lottery_worker_fee"`
	OracleContractAddr   string   `json:"oracle_contract_addr"`
	RegisterPrice        sdk.Uint `json:"register_price"`
	SafeLock             bool     `json:"safe_lock"`
	TotalCollect         sdk.Uint `json:"total_collect"`
	TotalFee             sdk.Uint `json:"total_fee"`
	TotalFeeWithdrew     sdk.Uint `json:"total_fee_withdrew"`
	TotalReward          sdk.Uint `json:"total_reward"`
	WinnerJackpotPercent uint8    `json:"winner_jackpot_percent"`
}

type RoundResponse struct {
	NextRound uint64 `json:"next_round"`
}
