package main

import (
	"context"
	"fmt"
	"os"

	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func main() {
	_, cosmosKeyring, err := chainclient.InitCosmosKeyring(
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

	address := "gio1fx8794synuvp9y8ft2w9rjdefpkj406ak8utpg"

	response, err := chainClient.GetFeegrant(context.Background(), address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tx Response: ", response.Allowances)
}
