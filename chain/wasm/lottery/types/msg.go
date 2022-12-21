package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

func NewMsgCollect(address *string, round *uint64) MsgCollect {
	return MsgCollect{
		Msg: Collect{
			Address: address,
			Round:   round,
		},
	}
}

func NewMsgLottery() MsgLottery {
	return MsgLottery{
		Msg: wasmtypes.Empty{},
	}
}

func NewMsgSafeLock() MsgSafeLock {
	return MsgSafeLock{
		Msg: wasmtypes.Empty{},
	}
}

func NewMsgSetConfig(
	blockTimePlay uint64,
	everyBlockTimePlay uint64,
	lotteryWorkerFee uint8,
	randContractAddress string,
	rankPercent uint8,
) MsgSetConfig {
	return MsgSetConfig{
		Msg: SetConfig{
			BlockTimePlay:       blockTimePlay,
			EveryBlockTimePlay:  everyBlockTimePlay,
			LotteryWorkerFee:    lotteryWorkerFee,
			RandContractAddress: randContractAddress,
			RankPercent:         rankPercent,
		},
	}
}

func NewMsgWithdrawFee(recipient string) MsgWithdrawFee {
	return MsgWithdrawFee{
		Msg: WithdrawFee{
			Recipient: recipient,
		},
	}
}

func NewQueryConfig() QueryConfig {
	return QueryConfig{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryCombination(address string, lotteryId uint64) QueryCombination {
	return QueryCombination{
		Query: Combination{
			Address:   address,
			LotteryId: lotteryId,
		},
	}
}

func NewQueryWinner(lotteryId uint64) QueryWinner {
	return QueryWinner{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryCountPlayer(lotteryId uint64) QueryCountPlayer {
	return QueryCountPlayer{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryCountTicket(lotteryId uint64) QueryCountTicket {
	return QueryCountTicket{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryCountWinner(lotteryId uint64, rank uint8) QueryCountWinner {
	return QueryCountWinner{
		Query: CountWinner{
			LotteryId: lotteryId,
			Rank:      rank,
		},
	}
}

func NewQueryWinningCombination(lotteryId uint64) QueryWinningCombination {
	return QueryWinningCombination{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryJackpot(lotteryId uint64) QueryJackpot {
	return QueryJackpot{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryPlayers(lotteryId uint64) QueryPlayers {
	return QueryPlayers{
		Query: ByLottery{
			LotteryId: lotteryId,
		},
	}
}

func NewQueryGetRound() QueryGetRound {
	return QueryGetRound{
		Query: wasmtypes.Empty{},
	}
}
