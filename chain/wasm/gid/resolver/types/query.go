package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type GetAddress struct {
	CoinType uint64 `json:"coin_type"`
	Node     []uint `json:"node"`
}

type QueryGetAddress struct {
	Query GetAddress `json:"get_address"`
}

type QueryGetTextData struct {
	Query TextData `json:"get_text_data"`
}

type GetAllTextData struct {
	Limit uint32 `json:"limit"`
	Node  []uint `json:"node"`
	Page  uint32 `json:"page"`
}

type QueryGetAllTextData struct {
	Query GetAllTextData `json:"get_all_text_data"`
}

type GetContentHash struct {
	Node []uint `json:"node"`
}

type QueryGetContentHash struct {
	Query GetContentHash `json:"get_content_hash"`
}

type QueryGetConfig struct {
	Query wasmtypes.Empty `json:"get_config"`
}
