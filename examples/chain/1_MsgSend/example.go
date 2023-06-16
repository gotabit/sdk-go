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
		"actual accuse plastic supply favorite banner trial company cloud wasp enable cactus",
		"",
		false,
	)

	if err != nil {
		panic(err)
	}

	chainClient, err := chainclient.NewChainClient(
		"devnet",
		cosmosKeyring,
		"1ugtb",
		nil,
	)

	if err != nil {
		panic(err)
	}

	// prepare tx msg
	msg := &banktypes.MsgSend{
		FromAddress: senderAddress.String(),
		ToAddress:   "gio1mh7rhytejck6xqxtcljls3pwj6wmeeqrjl2w2l",
		Amount: []sdktypes.Coin{{
			Denom: "ugtb", Amount: sdktypes.NewInt(1000000)}, // 1 ugtb
		},
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tx Hash: ", response.TxResponse.TxHash)

	fmt.Println("Tx Response: ", response.TxResponse)
}
