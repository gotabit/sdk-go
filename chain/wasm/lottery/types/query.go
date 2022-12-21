package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type QueryConfig struct {
	Query wasmtypes.Empty `json:"config"`
}

type Combination struct {
	Address   string `json:"string"`
	LotteryId uint64 `json:"lottery_id"`
}

type QueryCombination struct {
	Query Combination `json:"combination"`
}

type ByLottery struct {
	LotteryId uint64 `json:"lottery_id"`
}

type QueryWinner struct {
	Query ByLottery `json:"winner"`
}

type QueryCountPlayer struct {
	Query ByLottery `json:"count_player"`
}

type QueryCountTicket struct {
	Query ByLottery `json:"count_ticket"`
}

type CountWinner struct {
	LotteryId uint64 `json:"lottery_id"`
	Rank      uint8  `json:"rank"`
}

type QueryCountWinner struct {
	Query CountWinner `json:"count_winner"`
}

type QueryWinningCombination struct {
	Query ByLottery `json:"winning_combination"`
}

type QueryJackpot struct {
	Query ByLottery `json:"jackpot"`
}

type QueryPlayers struct {
	Query ByLottery `json:"players"`
}

type QueryGetRound struct {
	Query wasmtypes.Empty `json:"get_round"`
}
