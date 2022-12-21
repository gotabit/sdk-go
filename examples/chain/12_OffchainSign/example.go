package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	chainclient "github.com/gotabit/sdk-go/client/chain"
)

type Data struct {
	Address string `json:"address"`
	Gid     string `json:"gid"`
	Email   string `json:"email"`
	Time    int64  `json:"time"`
}

type unsafeExporter interface {
	// ExportPrivateKeyObject returns a private key in unarmored format.
	ExportPrivateKeyObject(uid string) (cryptotypes.PrivKey, error)
}

func main() {
	address, cosmosKeyring, err := chainclient.InitCosmosKeyring(
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

	pubkey, _ := hex.DecodeString("02925af9e4224a54b47c4f421c750ff0f5b1f098415b7b2c9a5612fffb52c76eb4")

	data := Data{Address: address.String(), Gid: "alice.gid", Email: "xxxx@gmail.com", Time: 1667639504}
	dataBytes, _ := json.Marshal(data)
	msg := string(dataBytes)

	sign, err := chainClient.OffChainSign(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sign: ", sign)

	isValid, err := chainClient.OffChainVerifyByPubkey(pubkey, msg, sign)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("pubkey verify result: ", isValid)

	isValid, err = chainClient.OffChainVerifyByAddress(address, msg, sign)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("address verify result: ", isValid)
}
