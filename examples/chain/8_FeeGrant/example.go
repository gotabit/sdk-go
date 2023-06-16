package main

import (
	"fmt"
	"os"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func main() {
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		chainclient.BaseCdc(),
		os.Getenv("HOME")+"/.gotabitd",
		"gotabitd",
		"file",
		"user",
		"",
		"climb cereal law remember october amount rough indicate trap gate slender moon",
		"",
		false,
	)

	if err != nil {
		panic(err)
	}

	granterAcc, err := sdktypes.AccAddressFromBech32("gio1fx8794synuvp9y8ft2w9rjdefpkj406ak8utpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	chainClient, err := chainclient.NewChainClient(
		"devnet",
		cosmosKeyring,
		"1ugtb",
		granterAcc,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// prepare tx msg
	msg := &banktypes.MsgSend{
		FromAddress: senderAddress.String(),
		ToAddress:   "gio1rmwwe2hge0lwzcu6ftmr547xsgjewqtdhehgte",
		Amount: []sdktypes.Coin{{
			Denom: "ugtb", Amount: sdktypes.NewInt(1000000)}, // 1 ugtb
		},
	}

	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tx Hash: ", response.TxResponse.TxHash)

	fmt.Println("Tx Response: ", response.TxResponse)
}
