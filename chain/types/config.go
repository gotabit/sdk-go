package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// GotabitBech32AddressPrefix defines the Bech32 prefix used for Address on the Gotabit Chain
	GotabitBech32AddressPrefix = "gio"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = GotabitBech32AddressPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = GotabitBech32AddressPrefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = GotabitBech32AddressPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = GotabitBech32AddressPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = GotabitBech32AddressPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = GotabitBech32AddressPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic

	CoinType = 118
)

// SetBech32Prefixes sets the global prefixes to be used when serializing addresses and public keys to Bech32 strings.
func SetBech32Prefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

// SetCoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetCoinType(config *sdk.Config) {
	config.SetCoinType(CoinType)
}
