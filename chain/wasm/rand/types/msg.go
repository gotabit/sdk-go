package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

func NewMsgSetBounty(round uint64) MsgSetBounty {
	return MsgSetBounty{
		Msg: SetBounty{
			Round: round,
		},
	}
}

func NewMsgAdd(previousSignature []uint, round uint64, signature []uint) MsgAdd {
	return MsgAdd{
		Msg: Add{
			PreviousSignature: previousSignature,
			Round:             round,
			Signature:         signature,
		},
	}
}

func NewQueryConfig() QueryConfig {
	return QueryConfig{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryGet(round uint64) QueryGet {
	return QueryGet{
		Query: Get{
			Round: round,
		},
	}
}

func NewQueryLatest() QueryLatest {
	return QueryLatest{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryBounties() QueryBounties {
	return QueryBounties{
		Query: wasmtypes.Empty{},
	}
}
