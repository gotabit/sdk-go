package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Config struct {
	BountyDenom string `json:"bounty_denom"`
	GenesisTime uint64 `json:"genesis_time"`
	PubKey      []uint `json:"pubkey"`
	RoundPeriod uint64 `json:"round_period"`
}

type Coin struct {
	Amount sdk.Uint `json:"amount"`
	Denom  string   `json:"denom"`
}

type Bounty struct {
	Amount []Coin `json:"amount"`
	Round  uint64 `json:"round"`
}

type BountiesResponse struct {
	Bounties []Bounty `json:"bounties"`
}

type GetResponse struct {
	Randomness []uint `json:"randomness"`
	Time       uint64 `json:"time"`
}

type LatestResponse struct {
	Randomness []uint `json:"randomness"`
	Round      uint64 `json:"round"`
	Time       uint64 `json:"time"`
}
