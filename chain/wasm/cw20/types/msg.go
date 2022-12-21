package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewCw20ReceiveMsg(amount sdk.Uint, msg []uint, sender string) Cw20ReceiveMsg {
	return Cw20ReceiveMsg{
		Amount: amount,
		Msg:    msg,
		Sender: sender,
	}
}

func NewCw20TransferMsg(recipient, amount string) Cw20TransferMsg {
	return Cw20TransferMsg{
		Msg: Cw20Transfer{
			Recipient: recipient,
			Amount:    amount,
		},
	}
}

func NewCw20BalanceQuery(address string) Cw20BalanceQuery {
	return Cw20BalanceQuery{
		Query: Cw20Balance{
			Address: address,
		},
	}
}
