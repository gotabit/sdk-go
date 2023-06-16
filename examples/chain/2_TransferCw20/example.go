package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cw20types "github.com/gotabit/sdk-go/chain/wasm/cw20/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func getBalance(chainClient chainclient.ChainClient, cw20Address, address string) string {
	query := cw20types.NewCw20BalanceQuery(address)
	queryData, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return "0"
	}

	ctx := context.Background()
	resp, err := chainClient.SmartContractState(ctx, cw20Address, queryData)
	if err != nil {
		fmt.Println(err)
		return "0"
	}
	res := cw20types.Cw20BalanceResponse{}
	if err = json.Unmarshal(resp.Data.Bytes(), &res); err != nil {
		fmt.Println(err)
		return "0"
	}

	return res.Balance
}

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
		panic(err)
	}

	cw20TokenAddress := "gio14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sl0hu5t"
	toAddress := "gio1r4rzuyt2a4rarlcqjcmzlk4rlea2rsd0e4d9z3"

	senderBeforeBalance := getBalance(chainClient, cw20TokenAddress, senderAddress.String())
	toBeforeBalance := getBalance(chainClient, cw20TokenAddress, toAddress)
	fmt.Println("from wallet balance (before transfer):", senderBeforeBalance)
	fmt.Println("to wallet balance (before transfer):", toBeforeBalance)

	// prepare tx msg
	transfer := cw20types.NewCw20TransferMsg(toAddress, "1")
	transferBz, err := json.Marshal(transfer)
	if err != nil {
		panic(err)
	}
	msg := &wasmtypes.MsgExecuteContract{
		Sender:   senderAddress.String(),
		Contract: cw20TokenAddress,
		Msg:      transferBz,
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tx Hash: ", response.TxResponse.TxHash)

	fmt.Println("Tx Response: ", response.TxResponse)

	senderBeforeBalance = getBalance(chainClient, cw20TokenAddress, senderAddress.String())
	toBeforeBalance = getBalance(chainClient, cw20TokenAddress, toAddress)
	fmt.Println("from wallet balance (after transfer):", senderBeforeBalance)
	fmt.Println("to wallet balance (after transfer):", toBeforeBalance)
}
