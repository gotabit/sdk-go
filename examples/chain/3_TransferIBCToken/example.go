package main

import (
	"fmt"
	"os"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func main() {
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
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
		fmt.Println(err)
		return
	}

	// prepare tx msg

	msg := &ibctransfertypes.MsgTransfer{
		SourcePort:    "transfer",
		SourceChannel: "", // source channel
		Token: sdktypes.Coin{
			Denom:  "ugtb",
			Amount: sdktypes.NewInt(1000000),
		},
		Sender:   senderAddress.String(),
		Receiver: "", // receiver address on dest chain
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
