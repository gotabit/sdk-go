package types

import (
	cw20types "github.com/gotabit/sdk-go/chain/wasm/cw20/types"
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type Collect struct {
	Address *string `json:"address,omitempty"`
	Round   *uint64 `json:"round,omitempty"`
}

type MsgCollect struct {
	Msg Collect `json:"collect"`
}

type MsgLottery struct {
	Msg wasmtypes.Empty `json:"lottery"`
}

type MsgSafeLock struct {
	Msg wasmtypes.Empty `json:"safe_lock"`
}

type MsgReceive struct {
	Msg cw20types.Cw20ReceiveMsg `json:"receive"`
}

type SetConfig struct {
	BlockTimePlay       uint64 `json:"block_time_play"`
	EveryBlockTimePlay  uint64 `json:"every_block_time_play"`
	LotteryWorkerFee    uint8  `json:"lottery_worker_fee"`
	RandContractAddress string `json:"rand_contract_address"`
	RankPercent         uint8  `json:"rank_percent"`
}

type MsgSetConfig struct {
	Msg SetConfig `json:"set_config"`
}

type WithdrawFee struct {
	Recipient string `json:"recipient"`
}

type MsgWithdrawFee struct {
	Msg WithdrawFee `json:"withdraw_fee"`
}
