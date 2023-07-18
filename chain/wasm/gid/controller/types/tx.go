package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

type Commit struct {
	Commitment  string    `json:"commitment"`
	Commitments *[]string `json:"commitments,omitempty"`
}

type MsgCommit struct {
	Msg Commit `json:"commit"`
}

type Register struct {
	Owner  string `json:"owner"`
	Domain Domain `json:"domain"`
}

type Domain struct {
	TopLevel TopLevel `json:"top_level"`
}

type TopLevel struct {
	Address      *string  `json:"address,omitempty"`
	Names        []string `json:"names"`
	DurationType string   `json:"duration_type"`
	Resolver     *string  `json:"resolver,omitempty"`
	Secret       string   `json:"secret"`
}

type MsgRegister struct {
	Msg Register `json:"register"`
}

type OwnerRegister struct {
	Address  *string `json:"address,omitempty"`
	Duration uint64  `json:"duration"`
	Name     string  `json:"name"`
	Owner    string  `json:"owner"`
	Resolver *string `json:"resolver,omitempty"`
}

type MsgOwnerRegister struct {
	Msg OwnerRegister `json:"owner_register"`
}

type SetConfig struct {
	EnableRegistration      bool   `json:"enable_registration"`
	L1                      uint32 `json:"l1"`
	L1LuckyNum              uint32 `json:"l1_lucky_num"`
	L1RegTime               uint64 `json:"l1_reg_time"`
	L2                      uint32 `json:"l2"`
	L2LuckyNum              uint32 `json:"l2_lucky_num"`
	L2RegTime               uint64 `json:"l2_reg_time"`
	L3                      uint32 `json:"l3"`
	L3LuckyNum              uint32 `json:"l3_lucky_num"`
	L3RegTime               uint64 `json:"l3_reg_time"`
	MaxCommitmentAge        uint64 `json:"max_commitment_age"`
	MinCommitmentAge        uint64 `json:"min_commitment_age"`
	MinRegistrationDuration uint64 `json:"min_registration_duration"`
	Owner                   string `json:"owner"`
	PermPrice               uint64 `json:"perm_price"`
	RegistrarAddress        string `json:"registrar_address"`
	Tier1Price              uint64 `json:"tier1_price"`
	Tier2Price              uint64 `json:"tier2_price"`
	Tier3Price              uint64 `json:"tier3_price"`
}

type MsgSetConfig struct {
	Msg SetConfig `json:"set_config"`
}

type MsgWithdraw struct {
	Msg wasmtypes.Empty `json:"withdraw"`
}

type Renew struct {
	Duration uint64 `json:"duration"`
	Name     string `json:"string"`
}

type MsgRenew struct {
	Msg Renew `json:"renew"`
}

type MsgOwnerRenew struct {
	Msg Renew `json:"owner_renew"`
}

type SetEnableRegistration struct {
	EnableRegistration bool `json:"enable_registration"`
}

type MsgSetEnableRegistration struct {
	Msg SetEnableRegistration `json:"set_enable_registration"`
}

type MsgAddReservedName struct {
	Msg Name `json:"add_reserved_name"`
}

type MsgRemoveReservedName struct {
	Msg Name `json:"remove_reserved_name"`
}

type ReverseClaim struct {
	Name     *string `json:"name,omitempty"`
	Owner    string  `json:"owner"`
	Resolver *string `json:"resolver,omitempty"`
}

type MsgReverseClaim struct {
	Msg ReverseClaim `json:"reverse_claim"`
}

type MsgNamePreview struct {
	Msg Name `json:"name_preview"`
}
