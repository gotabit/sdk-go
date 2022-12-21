<div align="center">

<a href="https://github.com/gotabit/sdk-ts"><img alt="GotaBit" src="https://res.gotabit.io/svg/icon.svg" width="150"/></a>

## GotaBit Golang SDK

</div>

## Installation

**Step 1: Install Golang**

Go v1.18+ or higher is required for the GotaBit Golang SDK.

1. Install [Go 1.18+ from the official site](https://go.dev/dl/). Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Linux:

   ```sh
   wget <https://golang.org/dl/go1.18.2.linux-amd64.tar.gz>
   sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Confirm your Go installation by checking the version:

   ```sh
   go version
   ```


**Step 2: Get source code**

Clone the repository locally and install needed dependencies

```bash
$ git clone git@github.com:hjcore/sdk-go.git
$ cd sdk-go
$ go install ./...
```

## Documentation

Check the [documentation](docs.md) to get you started!

## Run examples
```bash
# run chain example
go run examples/chain/1_MsgSend/example.go

# run stake example
go run examples/chain/4_Stake/example.go
```

## Talk to us

We have active, helpful communities on Twitter, Discord, and Telegram.

<p>
<a href="https://twitter.com/GotaBitG"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" alt="Tweet" height="30"/></a> 
  &nbsp;
 <a href="https://t.me/GotaBitG"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram" height="30"/></a> 
</p>

For updates on the GotaBit team's activities follow us on the [GotaBit Twitter](https://twitter.com/GotaBitG) account.

## License

This software is licensed under the Apache 2.0 license.

© 2022 GotaBit Limited
