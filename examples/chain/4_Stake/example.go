package main

import (
	"fmt"
	"os"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
		"dinner proud piano mention silk plunge forest fold trial duck electric define",
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
	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: senderAddress.String(),
		ValidatorAddress: "giovaloper1qdgzfy4vta5p43l4urdtmawka3qv2ldhvh88jv",
		Amount: sdktypes.Coin{
			Denom: "ugtb", Amount: sdktypes.NewInt(10000000), // 10 ugtb
		},
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
