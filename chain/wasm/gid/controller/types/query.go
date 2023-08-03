package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type QueryOwner struct {
	Query wasmtypes.Empty `json:"owner"`
}

type QueryRegistrar struct {
	Query wasmtypes.Empty `json:"registrar"`
}

type CommitmentTimestamp struct {
	Commitment string `json:"commitment"`
}

type QueryCommitmentTimestamp struct {
	Query CommitmentTimestamp `json:"commitment_timestamp"`
}

type GetCommitment struct {
	Address  *string `json:"address,omitempty"`
	Name     string  `json:"name"`
	Owner    string  `json:"owner"`
	Resolver *string `json:"resolver,omitempty"`
	Secret   string  `json:"secret"`
}

type QueryGetCommitment struct {
	Query GetCommitment `json:"get_commitment"`
}

type RentPrice struct {
	DurationType string `json:"duration_type"`
	Name         string `json:"name"`
}

type QueryRentPrice struct {
	Query RentPrice `json:"rent_price"`
}

type QueryMaxCommitmentAge struct {
	Query wasmtypes.Empty `json:"max_commitment_age"`
}

type QueryMinCommitmentAge struct {
	Query wasmtypes.Empty `json:"min_commitment_age"`
}

type QueryMinRegistrationDuration struct {
	Query wasmtypes.Empty `json:"min_registration_duration"`
}

type QueryIsValidName struct {
	Query Name `json:"is_valid_name"`
}

type QueryGetTokenId struct {
	Query Name `json:"get_token_id"`
}

type QueryGetNodeHash struct {
	Query Name `json:"get_nodehash"`
}

type QueryGetNodeInfo struct {
	Query Name `json:"get_node_info"`
}

type QueryGetPrice struct {
	Query wasmtypes.Empty `json:"get_price"`
}

type QueryGetReverseNode struct {
	Query Address `json:"get_reverse_node"`
}

type QueryGetRecord struct {
	Query Name `json:"get_record"`
}

type QueryNamePreview struct {
	IsAvailable bool `json:"is_available"`
	IsValidName bool `json:"is_valid_name"`
	IsReserved  bool `json:"is_reserved"`
}
