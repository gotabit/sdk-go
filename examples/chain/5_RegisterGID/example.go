package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	controllertypes "github.com/gotabit/sdk-go/chain/wasm/gid/controller/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func GetRentPrice(chainClient chainclient.ChainClient, controller string, durationType, name string) controllertypes.RentPriceResponse {
	query := controllertypes.NewQueryRentPrice(durationType, name)
	queryData, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return controllertypes.RentPriceResponse{}
	}

	ctx := context.Background()
	resp, err := chainClient.SmartContractState(ctx, controller, queryData)
	if err != nil {
		fmt.Println(err)
		return controllertypes.RentPriceResponse{}
	}
	res := controllertypes.RentPriceResponse{}
	if err = json.Unmarshal(resp.Data.Bytes(), &res); err != nil {
		fmt.Println(err)
		return controllertypes.RentPriceResponse{}
	}

	return res
}

func GetCommitment(chainClient chainclient.ChainClient, controller, resolver, owner, name, secret string) controllertypes.GetCommitmentResponse {
	query := controllertypes.NewQueryGetCommitment(nil, name, owner, &resolver, secret)
	queryData, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return controllertypes.GetCommitmentResponse{}
	}

	ctx := context.Background()
	resp, err := chainClient.SmartContractState(ctx, controller, queryData)
	if err != nil {
		fmt.Println(err)
		return controllertypes.GetCommitmentResponse{}
	}
	res := controllertypes.GetCommitmentResponse{}
	if err = json.Unmarshal(resp.Data.Bytes(), &res); err != nil {
		fmt.Println(err)
		return controllertypes.GetCommitmentResponse{}
	}

	return res
}

func main() {
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		chainclient.BaseCdc(),
		os.Getenv("HOME")+"/.gotabitd",
		"gotabitd",
		"file",
		"user",
		"",
		"sign lonely broken town organ capable furnace hip leaf gospel engage clay",
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

	name := "alice1"
	owner := senderAddress.String()
	secret := "11111111"
	durationType := "year"

	rentPrice := GetRentPrice(chainClient, controllerAddress, durationType, name)
	commitment := GetCommitment(chainClient, controllerAddress, resolverAddress, owner, name, secret)

	// prepare tx msg
	commitCommitmentMsg := controllertypes.NewMsgCommit(commitment.Commitment)
	bz, err := json.Marshal(commitCommitmentMsg)
	if err != nil {
		panic(err)
	}
	msg := &wasmtypes.MsgExecuteContract{
		Sender:   senderAddress.String(),
		Contract: controllerAddress,
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

	time.Sleep(time.Second * 7)

	// prepare tx msg
	registerMsg := controllertypes.NewMsgRegister(nil, durationType, name, owner, &resolverAddress, secret)
	registerMsgBz, err := json.Marshal(registerMsg)
	if err != nil {
		panic(err)
	}
	msg = &wasmtypes.MsgExecuteContract{
		Sender:   senderAddress.String(),
		Contract: controllerAddress,
		Msg:      registerMsgBz,
		Funds: []sdktypes.Coin{{
			Denom: "ugtb", Amount: sdktypes.NewInt(rentPrice.Price.BigInt().Int64())},
		},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response2, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tx Hash: ", response2.TxResponse.TxHash)

	fmt.Println("Tx Response: ", response2.TxResponse)
}
