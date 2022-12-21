package main

import (
	"fmt"
	"time"

	chainclient "github.com/gotabit/sdk-go/client/chain"
)

func main() {
	mnemonic, account, cosmosKeyring, err := chainclient.CreateMnemonicAndKeyring(12, "user", "")
	fmt.Println("Generated Mnemonic:", mnemonic)

	if err != nil {
		panic(err)
	}

	t := time.Now()
	filename := fmt.Sprintf("backup/backup-%s.key", t.Format("20060102150405"))
	err = chainclient.ExportAccount(filename, cosmosKeyring, account, "passphrase")
	if err != nil {
		panic(err)
	}
}
