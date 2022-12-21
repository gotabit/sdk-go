package chain

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	ethereumsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type unsafeExporter interface {
	// ExportPrivateKeyObject returns a private key in unarmored format.
	ExportPrivateKeyObject(uid string) (cryptotypes.PrivKey, error)
}

func (c *chainClient) OffChainSign(
	msg string,
) (string, error) {
	privKey, err := c.ctx.Keyring.(unsafeExporter).ExportPrivateKeyObject(defaultKeyringKeyName)
	if err != nil {
		return "", err
	}

	msg = c.buildMsg(msg, c.ctx.FromAddress)

	priv := secp256k1.PrivKey{Key: privKey.Bytes()}

	sign, err := priv.Sign([]byte(msg))
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.WithPadding('=').EncodeToString(sign), nil
}

func (c *chainClient) OffChainVerifyByPubkey(
	pubkey []byte,
	msg string,
	sign string,
) (bool, error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	pub := &secp256k1.PubKey{Key: pubkey}

	msg = c.buildMsg(msg, sdk.AccAddress(pub.Address()))

	isValid := pub.VerifySignature([]byte(msg), signBytes[:64])
	return isValid, nil
}

func (c *chainClient) OffChainVerifyByAddress(
	address sdk.AccAddress,
	msg string,
	sign string,
) (bool, error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	msg = c.buildMsg(msg, address)

	newSignBytes := make([]byte, 65)
	copy(newSignBytes, signBytes)

	sum := sha256.Sum256([]byte(msg))

	for i := 0; i < 4; i++ {
		newSignBytes[64] = byte(i)
		pubkey, err := ethereumsecp256k1.RecoverPubkey(sum[:], newSignBytes)
		if err != nil {
			continue
		}

		x := new(big.Int).SetBytes(pubkey[1:33])
		y := new(big.Int).SetBytes(pubkey[33:])
		pubkey = ethereumsecp256k1.CompressPubkey(x, y)

		if len(pubkey) != 33 {
			continue
		}

		pub := &secp256k1.PubKey{Key: pubkey}

		convertAddress, err := bech32.ConvertAndEncode("gio", pub.Address().Bytes())
		if err != nil {
			continue
		}

		if convertAddress == address.String() {
			return true, nil
		}
	}
	return false, nil
}

func (c *chainClient) buildMsg(data string, address sdk.AccAddress) string {
	data = fmt.Sprintf(`{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"%s","signer":"%s"}}],"sequence":"0"}`,
		base64.RawStdEncoding.WithPadding('=').EncodeToString([]byte(data)), address.String())

	return data
}
