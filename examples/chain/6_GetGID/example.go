package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

func GetTextResult(chainClient chainclient.ChainClient, resolver string, node []uint, key string) resolvertypes.TextDataResponse {
	query := resolvertypes.NewQueryGetTextData(key, node)
	queryData, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return resolvertypes.TextDataResponse{}
	}

	ctx := context.Background()
	resp, err := chainClient.SmartContractState(ctx, resolver, queryData)
	if err != nil {
		fmt.Println(err)
		return resolvertypes.TextDataResponse{}
	}

	res := resolvertypes.TextDataResponse{}
	if err = json.Unmarshal(resp.Data.Bytes(), &res); err != nil {
		fmt.Println(err)
		return resolvertypes.TextDataResponse{}
	}

	return res
}

func main() {
	_, cosmosKeyring, err := chainclient.InitCosmosKeyring(
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

	resolverAddress := "gio1u45dw8dsvlnuq5rq9jaasy2glmwf265ux29smtt8ecuf9e7rkmcsl73t52"
	controllerAddress := "gio1m8pkv82eamkvw84w38auduy9r5kqp3hf8gwtykwk6d38augx6dzsgcsg9c"

	nodehash := GetNodeHash(chainClient, controllerAddress, "alice101")
	fmt.Println("nodehash:", nodehash)
	name := GetTextResult(chainClient, resolverAddress, nodehash.Node, "name")
	fmt.Println("get_text_result:", name)
}
