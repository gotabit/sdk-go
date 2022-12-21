package main

import (
	"fmt"
	"os"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant"
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

	granteeAcc, err := sdktypes.AccAddressFromBech32("gio1fx8794synuvp9y8ft2w9rjdefpkj406ak8utpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	allowance := feegrant.BasicAllowance{
		SpendLimit: []sdktypes.Coin{{
			Denom: "ugtb", Amount: sdktypes.NewInt(100000000)}, // 100 ugtb
		},
	}

	msg, err := feegrant.NewMsgGrantAllowance(feegrant.FeeAllowanceI(&allowance), senderAddress, granteeAcc)
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
