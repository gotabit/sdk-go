package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type QueryConfig struct {
	Query wasmtypes.Empty `json:"config"`
}

type Get struct {
	Round uint64 `json:"round"`
}

type QueryGet struct {
	Query Get `json:"get"`
}

type QueryLatest struct {
	Query wasmtypes.Empty `json:"latest"`
}

type QueryBounties struct {
	Query wasmtypes.Empty `json:"bounties"`
}
