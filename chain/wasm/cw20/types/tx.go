package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Cw20ReceiveMsg struct {
	Amount sdk.Uint `json:"amount"`
	Msg    []uint   `json:"msg"`
	Sender string   `json:"sender"`
}

type Cw20Transfer struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type Cw20TransferMsg struct {
	Msg Cw20Transfer `json:"transfer"`
}
