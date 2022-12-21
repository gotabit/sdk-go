package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	controllertypes "github.com/gotabit/sdk-go/chain/wasm/gid/controller/types"
	resolvertypes "github.com/gotabit/sdk-go/chain/wasm/gid/resolver/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func GetNodeHash(chainClient chainclient.ChainClient, controller, name string) controllertypes.NodeHashResponse {
	query := controllertypes.NewQueryGetNodeHash(name)
	queryData, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return controllertypes.NodeHashResponse{}
	}

	ctx := context.Background()
	resp, err := chainClient.SmartContractState(ctx, controller, queryData)
	if err != nil {
		fmt.Println(err)
		return controllertypes.NodeHashResponse{}
	}
	res := controllertypes.NodeHashResponse{}
	if err = json.Unmarshal(resp.Data.Bytes(), &res); err != nil {
		fmt.Println(err)
		return controllertypes.NodeHashResponse{}
	}

	return res
}

func main() {
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
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

	resolverAddress := "gio1eyfccmjm6732k7wp4p6gdjwhxjwsvje44j0hfx8nkgrm8fs7vqfsvq7yce"
	controllerAddress := "gio1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqhhpzlg"

	nodehash := GetNodeHash(chainClient, controllerAddress, "alice101")
	fmt.Println("nodehash:", nodehash)

	// prepare tx msg
	setTextDataMsg := resolvertypes.NewMsgSetTextData("name", nodehash.Node, "alice001")
	bz, err := json.Marshal(setTextDataMsg)
	if err != nil {
		panic(err)
	}
	msg := &wasmtypes.MsgExecuteContract{
		Sender:   senderAddress.String(),
		Contract: resolverAddress,
		Msg:      bz,
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
