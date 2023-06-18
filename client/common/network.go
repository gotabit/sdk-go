package common

import (
	"context"
	"net"
	"strings"
)

type Network struct {
	RPCEndpoint  string
	GrpcEndpoint string
	ChainId      string
	Name         string
	RestEndpoint string
	Denom        string
	Decimals     int
}

var Networks map[string]Network = map[string]Network{
	"mainnet": {
		RPCEndpoint:  "https://rpc.gotabit.dev:443",
		GrpcEndpoint: "https://grpc.gotabit.dev:443",
		ChainId:      "gotabit-alpha",
		Name:         "GotaBit",
		RestEndpoint: "https://rest.gotabit.dev:443",
		Denom:        "ugtb",
		Decimals:     6,
	},
	"testnet": {
		RPCEndpoint:  "https://rpc-testnet.gotabit.dev:443",
		GrpcEndpoint: "https://grpc-testnet.gotabit.dev:443",
		ChainId:      "gotabit-test-1",
		Name:         "GotaBit-test",
		RestEndpoint: "https://rest-testnet.gotabit.dev:443",
		Denom:        "ugtb",
		Decimals:     6,
	},
	"devnet": {
		RPCEndpoint:  "https://rpc-devnet.gotabit.dev:443",
		GrpcEndpoint: "https://grpc-devnet.gotabit.dev:443",
		ChainId:      "gotabit-dev-1",
		Name:         "GotaBit-dev",
		RestEndpoint: "https://api-devnet.gotabit.dev:443",
		Denom:        "ugtb",
		Decimals:     6,
	},
	"newsdk": {
		RPCEndpoint:  "https://rpc-testnet.gotabit.dev:443",
		GrpcEndpoint: "https://grpc-testnet.gotabit.dev:443",
		ChainId:      "gotabit-test-1",
		Name:         "GotaBit-test",
		RestEndpoint: "https://rest-testnet.gotabit.dev:443",
		Denom:        "ugtb",
		Decimals:     6,
	},
}

func LoadNetwork(name string) Network {
	return Networks[name]
}

func DialerFunc(ctx context.Context, addr string) (net.Conn, error) {
	return Connect(addr)
}

// Connect dials the given address and returns a net.Conn. The protoAddr argument should be prefixed with the protocol,
// eg. "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock"
func Connect(protoAddr string) (net.Conn, error) {
	proto, address := ProtocolAndAddress(protoAddr)
	conn, err := net.Dial(proto, address)
	return conn, err
}

// ProtocolAndAddress splits an address into the protocol and address components.
// For instance, "tcp://127.0.0.1:8080" will be split into "tcp" and "127.0.0.1:8080".
// If the address has no protocol prefix, the default is "tcp".
func ProtocolAndAddress(listenAddr string) (string, string) {
	protocol, address := "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
