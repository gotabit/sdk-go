package types

import (
	wasmtypes "github.com/gotabit/sdk-go/chain/wasm/types"
)

func NewMsgSetAddress(address string, coinType uint64, node []uint) MsgSetAddress {
	return MsgSetAddress{
		Msg: SetAddress{
			Address:  address,
			CoinType: coinType,
			Node:     node,
		},
	}
}

func NewMsgSetTextData(key string, node []uint, value string) MsgSetTextData {
	return MsgSetTextData{
		Msg: SetTextData{
			Key:   key,
			Node:  node,
			Value: value,
		},
	}
}

func NewMsgSetContentHash(hash, node []uint) MsgSetContentHash {
	return MsgSetContentHash{
		Msg: SetContentHash{
			Hash: hash,
			Node: node,
		},
	}
}

func NewMsgSetConfig(interfaceId uint64, owner, registryAddress string) MsgSetConfig {
	return MsgSetConfig{
		Msg: Config{
			IntefaceId:      interfaceId,
			Owner:           owner,
			RegistryAddress: registryAddress,
		},
	}
}

func NewMsgDeleteTextData(key string, node []uint) MsgDeleteTextData {
	return MsgDeleteTextData{
		Msg: TextData{
			Key:  key,
			Node: node,
		},
	}
}

func NewQueryGetAddress(coinType uint64, node []uint) QueryGetAddress {
	return QueryGetAddress{
		Query: GetAddress{
			CoinType: coinType,
			Node:     node,
		},
	}
}

func NewQueryGetTextData(key string, node []uint) QueryGetTextData {
	return QueryGetTextData{
		Query: TextData{
			Key:  key,
			Node: node,
		},
	}
}

func NewQueryGetAllTextData(limit uint32, node []uint, page uint32) QueryGetAllTextData {
	return QueryGetAllTextData{
		Query: GetAllTextData{
			Limit: limit,
			Node:  node,
			Page:  page,
		},
	}
}

func NewQueryGetContentHash(node []uint) QueryGetContentHash {
	return QueryGetContentHash{
		Query: GetContentHash{
			Node: node,
		},
	}
}

func NewQueryGetConfig() QueryGetConfig {
	return QueryGetConfig{
		Query: wasmtypes.Empty{},
	}
}
