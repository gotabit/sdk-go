# GotaBit Golang SDK Documenatation

## How to use SDK

This SDK consists of many useful functions and types that can be used to interact with GotaBit mainnet and testnet.

It can also interact with CosmWasm smart contracts deployed on Gotabit.

Any ideas or suggestions are welcome!

## How to make example

### Step 1. Load Network & Setup RPC

```code
import (
    "github.com/gotabit/sdk-go/client/common"
    rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

network := common.LoadNetwork("testnet")
rpc, err := rpchttp.New(network.RPCEndpoint, "/websocket")
```

### Step 2. Import/Create Wallet and Keyring

```code
import (
    chainclient "github.com/gotabit/sdk-go/client/chain"
)

address, cosmosKeyring, err := chainclient.InitCosmosKeyring(
    os.Getenv("HOME")+"/.gotabitd",                                 // dir
    "gotabitd",                                                     // appname
    "file",                                                         // backend
    "user",                                                         // key name
    "",                                                             // passphrase
    "test test test test test test test test test test test test",  // mnemonic
    "",                                                             // private key
    false,                                                          // use ledger
)
```

### Step 3. Create Client Context

```code
import (
    chainclient "github.com/gotabit/sdk-go/client/chain"
)

clientCtx, err := chainclient.NewClientContext(
    network.ChainId,
    senderAddress.String(),
    cosmosKeyring,
)

clientCtx = clientCtx.WithNodeURI(network.RPCEndpoint).WithClient(rpc)
```

### Step 4. Create Chain Client

```code
import (
    "github.com/gotabit/sdk-go/client/common"
    chainclient "github.com/gotabit/sdk-go/client/chain"
)

chainClient, err := chainclient.NewChainClient(
    clientCtx,
    network.GrpcEndpoint,
    common.OptionGasPrices("1ugtb"),
	common.OptionTLSCert(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})),
)
```

### Step 5. Prepare Tx Messages

```code
import (
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

msg := &banktypes.MsgSend{
    FromAddress: address.String(),
    ToAddress:   "gio1rmwwe2hge0lwzcu6ftmr547xsgjewqtdhehgte",
    Amount: []sdktypes.Coin{{
        Denom: "ugtb", Amount: sdktypes.NewInt(1000000)}, // 1 ugtb
    },
}
```

### Step 6. Broadcat Message

```code
response, err := chainClient.SyncBroadcastMsg(msg)
```
