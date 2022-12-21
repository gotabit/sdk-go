package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GotabitCoin string = "ugtb"
)

// NewGotabitCoin is a utility function that returns an "ugtb" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewGotabitCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(GotabitCoin, amount)
}

// NewGotabitDecCoin is a utility function that returns an "ugtb" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewGotabitDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(GotabitCoin, amount)
}

// NewGotabitCoinInt64 is a utility function that returns an "ugtb" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewGotabitCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(GotabitCoin, amount)
}
