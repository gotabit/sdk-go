package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

func NewMsgCommit(commitment string) MsgCommit {
	return MsgCommit{
		Msg: Commit{
			Commitment: commitment,
		},
	}
}

func NewMsgRegister(
	address *string,
	durationType string,
	names []string,
	owner string,
	resolver *string,
	secret string,
) MsgRegister {
	return MsgRegister{
		Msg: Register{
			Owner: owner,
			Domain: Domain{
				TopLevel: TopLevel{
					Address:      address,
					Names:        names,
					DurationType: durationType,
					Resolver:     resolver,
					Secret:       secret,
				},
			},
		},
	}
}

func NewMsgOwnerRegister(
	address *string,
	duration uint64,
	name string,
	owner string,
	resolver *string,
) MsgOwnerRegister {
	return MsgOwnerRegister{
		Msg: OwnerRegister{
			Address:  address,
			Duration: duration,
			Name:     name,
			Owner:    owner,
			Resolver: resolver,
		},
	}
}

func NewMsgSetConfig(
	enableRegistration bool,
	l1 uint32,
	l1LuckyNum uint32,
	l1RegTime uint64,
	l2 uint32,
	l2LuckyNum uint32,
	l2RegTime uint64,
	l3 uint32,
	l3LuckyNum uint32,
	l3RegTime uint64,
	maxCommitmentAge uint64,
	minCommitmentAge uint64,
	minRegistrationDuration uint64,
	owner string,
	permPrice uint64,
	registrarAddress string,
	tier1Price uint64,
	tier2Price uint64,
	tier3Price uint64,
) MsgSetConfig {
	return MsgSetConfig{
		Msg: SetConfig{
			EnableRegistration:      enableRegistration,
			L1:                      l1,
			L1LuckyNum:              l1LuckyNum,
			L1RegTime:               l1RegTime,
			L2:                      l2,
			L2LuckyNum:              l2LuckyNum,
			L2RegTime:               l2RegTime,
			L3:                      l3,
			L3LuckyNum:              l3LuckyNum,
			L3RegTime:               l3RegTime,
			MaxCommitmentAge:        maxCommitmentAge,
			MinCommitmentAge:        minCommitmentAge,
			MinRegistrationDuration: minRegistrationDuration,
			Owner:                   owner,
			PermPrice:               permPrice,
			RegistrarAddress:        registrarAddress,
			Tier1Price:              tier1Price,
			Tier2Price:              tier2Price,
			Tier3Price:              tier3Price,
		},
	}
}

func NewMsgWithdraw() MsgWithdraw {
	return MsgWithdraw{
		Msg: wasmtypes.Empty{},
	}
}

func NewMsgRenew(duration uint64, name string) MsgRenew {
	return MsgRenew{
		Msg: Renew{
			Duration: duration,
			Name:     name,
		},
	}
}

func NewMsgOwnerRenew(duration uint64, name string) MsgOwnerRenew {
	return MsgOwnerRenew{
		Msg: Renew{
			Duration: duration,
			Name:     name,
		},
	}
}

func NewMsgSetEnableRegistration(enableRegistration bool) MsgSetEnableRegistration {
	return MsgSetEnableRegistration{
		Msg: SetEnableRegistration{
			EnableRegistration: enableRegistration,
		},
	}
}

func NewMsgAddReservedName(name string) MsgAddReservedName {
	return MsgAddReservedName{
		Msg: Name{
			Name: name,
		},
	}
}

func NewMsgRemoveReservedName(name string) MsgRemoveReservedName {
	return MsgRemoveReservedName{
		Msg: Name{
			Name: name,
		},
	}
}

func NewMsgReverseClaim(name *string, owner string, resolver *string) MsgReverseClaim {
	return MsgReverseClaim{
		Msg: ReverseClaim{
			Name:     name,
			Owner:    owner,
			Resolver: resolver,
		},
	}
}

func NewQueryOwner() QueryOwner {
	return QueryOwner{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryRegistrar() QueryRegistrar {
	return QueryRegistrar{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryCommitmentTimestamp(commitment string) QueryCommitmentTimestamp {
	return QueryCommitmentTimestamp{
		Query: CommitmentTimestamp{
			Commitment: commitment,
		},
	}
}

func NewQueryGetCommitment(
	address *string,
	name string,
	owner string,
	resolver *string,
	secret string,
) QueryGetCommitment {
	return QueryGetCommitment{
		Query: GetCommitment{
			Address:  address,
			Name:     name,
			Owner:    owner,
			Resolver: resolver,
			Secret:   secret,
		},
	}
}

func NewQueryRentPrice(durationType, name string) QueryRentPrice {
	return QueryRentPrice{
		Query: RentPrice{
			DurationType: durationType,
			Name:         name,
		},
	}
}

func NewQueryMaxCommitmentAge() QueryMaxCommitmentAge {
	return QueryMaxCommitmentAge{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryMinCommitmentAge() QueryMinCommitmentAge {
	return QueryMinCommitmentAge{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryMinRegistrationDuration() QueryMinRegistrationDuration {
	return QueryMinRegistrationDuration{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryIsValidName(name string) QueryIsValidName {
	return QueryIsValidName{
		Query: Name{
			Name: name,
		},
	}
}

func NewQueryGetTokenId(name string) QueryGetTokenId {
	return QueryGetTokenId{
		Query: Name{
			Name: name,
		},
	}
}

func NewQueryGetNodeHash(name string) QueryGetNodeHash {
	return QueryGetNodeHash{
		Query: Name{
			Name: name,
		},
	}
}

func NewQueryGetNodeInfo(name string) QueryGetNodeInfo {
	return QueryGetNodeInfo{
		Query: Name{
			Name: name,
		},
	}
}

func NewQueryGetPrice() QueryGetPrice {
	return QueryGetPrice{
		Query: wasmtypes.Empty{},
	}
}

func NewQueryGetReverseNode(address string) QueryGetReverseNode {
	return QueryGetReverseNode{
		Query: Address{
			Address: address,
		},
	}
}

func NewNamePreview(name string) MsgNamePreview {
	return MsgNamePreview{
		Msg: Name{
			Name: name,
		},
	}
}
