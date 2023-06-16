package chain

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmcrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"
	"github.com/gotabit/sdk-go/client/common"
	"github.com/pkg/errors"
)

const defaultKeyringKeyName = "testuser"

var emptyCosmosAddress = cosmtypes.AccAddress{}

func InitCosmosKeyring(
	cdc codec.Codec,
	cosmosKeyringDir string,
	cosmosKeyringAppName string,
	cosmosKeyringBackend string,
	cosmosKeyFrom string,
	cosmosKeyPassphrase string,
	cosmosMnemonic string,
	cosmosPrivKey string,
	cosmosUseLedger bool,
) (cosmtypes.AccAddress, keyring.Keyring, error) {
	switch {
	case len(cosmosPrivKey) > 0:
		if cosmosUseLedger {
			err := errors.New("cannot combine ledger and privkey options")
			return emptyCosmosAddress, nil, err
		}

		pkBytes, err := common.HexToBytes(cosmosPrivKey)
		if err != nil {
			err = errors.Wrap(err, "failed to hex-decode cosmos account privkey")
			return emptyCosmosAddress, nil, err
		}

		cosmosAccPk := &secp256k1.PrivKey{
			Key: pkBytes,
		}

		addressFromPk := cosmtypes.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

		var keyName string

		// check that if cosmos 'From' specified separately, it must match the provided privkey,
		if len(cosmosKeyFrom) > 0 {
			addressFrom, err := cosmtypes.AccAddressFromBech32(cosmosKeyFrom)
			if err == nil {
				if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
					err = errors.Errorf("expected account address %s but got %s from the private key", addressFrom.String(), addressFromPk.String())
					return emptyCosmosAddress, nil, err
				}
			} else {
				// use it as a name then
				keyName = cosmosKeyFrom
			}
		}

		if len(keyName) == 0 {
			keyName = defaultKeyringKeyName
		}

		// wrap a PK into a Keyring
		kb, err := KeyringForPrivKey(cdc, keyName, cosmosAccPk)
		return addressFromPk, kb, err

	case len(cosmosMnemonic) > 0:
		var keyName string

		if len(keyName) == 0 {
			keyName = defaultKeyringKeyName
		}

		return KeyringForMnemonic(cdc, keyName, cosmosMnemonic)

	case len(cosmosKeyFrom) > 0:
		var fromIsAddress bool
		addressFrom, err := cosmtypes.AccAddressFromBech32(cosmosKeyFrom)
		if err == nil {
			fromIsAddress = true
		}

		var passReader io.Reader = os.Stdin
		if len(cosmosKeyPassphrase) > 0 {
			passReader = newPassReader(cosmosKeyPassphrase)
		}

		var absoluteKeyringDir string
		if filepath.IsAbs(cosmosKeyringDir) {
			absoluteKeyringDir = cosmosKeyringDir
		} else {
			absoluteKeyringDir, _ = filepath.Abs(cosmosKeyringDir)
		}

		kb, err := keyring.New(
			cosmosKeyringAppName,
			cosmosKeyringBackend,
			absoluteKeyringDir,
			passReader,
			cdc,
		)
		if err != nil {
			err = errors.Wrap(err, "failed to init keyring")
			return emptyCosmosAddress, nil, err
		}

		var keyInfo *keyring.Record
		if fromIsAddress {
			if keyInfo, err = kb.KeyByAddress(addressFrom); err != nil {
				err = errors.Wrapf(err, "couldn't find an entry for the key %s in keybase", addressFrom.String())
				return emptyCosmosAddress, nil, err
			}
		} else {
			if keyInfo, err = kb.Key(cosmosKeyFrom); err != nil {
				err = errors.Wrapf(err, "could not find an entry for the key '%s' in keybase", cosmosKeyFrom)
				return emptyCosmosAddress, nil, err
			}
		}

		addr, err := keyInfo.GetAddress()
		if err != nil {
			panic(err)
		}
		switch keyType := keyInfo.GetType(); keyType {
		case keyring.TypeLocal:
			// kb has a key and it's totally usable
			return addr, kb, nil
		case keyring.TypeLedger:
			// the kb stores references to ledger keys, so we must explicitly
			// check that. kb doesn't know how to scan HD keys - they must be added manually before
			if cosmosUseLedger {
				return addr, kb, nil
			}
			err := errors.Errorf("'%s' key is a ledger reference, enable ledger option", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		case keyring.TypeOffline:
			err := errors.Errorf("'%s' key is an offline key, not supported yet", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		case keyring.TypeMulti:
			err := errors.Errorf("'%s' key is an multisig key, not supported yet", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		default:
			err := errors.Errorf("'%s' key  has unsupported type: %s", keyInfo.Name, keyType)
			return emptyCosmosAddress, nil, err
		}

	default:
		err := errors.New("insufficient cosmos key details provided")
		return emptyCosmosAddress, nil, err
	}
}

func CreateMnemonicAndKeyring(cdc codec.Codec, length int, name string, passphrase string) (string, cosmtypes.AccAddress, keyring.Keyring, error) {
	if length != 12 && length != 15 && length != 18 && length != 21 && length != 24 {
		return "", emptyCosmosAddress, nil, errors.New("invalid mnemonic length")
	}
	entropy, err := bip39.NewEntropy(256 * length / 24)
	if err != nil {
		return "", emptyCosmosAddress, nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", emptyCosmosAddress, nil, err
	}

	kb := keyring.NewInMemory(cdc)
	info, err := kb.NewAccount(name, mnemonic, passphrase, "m/44'/118'/0'/0/0", hd.Secp256k1)
	if err != nil {
		err = errors.Wrap(err, "failed to create new mnemonic")
		return "", emptyCosmosAddress, nil, err
	}
	addr, err := info.GetAddress()
	if err != nil {
		return "", nil, nil, err
	}
	return mnemonic, addr, kb, nil
}

func ExportAccount(outputFile string, kb keyring.Keyring, account cosmtypes.AccAddress, encryptPassphrase string) error {
	armor, err := kb.ExportPrivKeyArmorByAddress(account, encryptPassphrase)
	if err != nil {
		return err
	}

	fmt.Println(armor)

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(armor)
	return err
}

func newPassReader(pass string) io.Reader {
	return &passReader{
		pass: pass,
		buf:  new(bytes.Buffer),
	}
}

type passReader struct {
	pass string
	buf  *bytes.Buffer
}

var _ io.Reader = &passReader{}

func (r *passReader) Read(p []byte) (n int, err error) {
	n, err = r.buf.Read(p)
	if err == io.EOF || n == 0 {
		r.buf.WriteString(r.pass + "\n")

		n, err = r.buf.Read(p)
	}

	return
}

// KeyringForPrivKey creates a temporary in-mem keyring for a PrivKey.
// Allows to init Context when the key has been provided in plaintext and parsed.
func KeyringForPrivKey(cdc codec.Codec, name string, privKey cryptotypes.PrivKey) (keyring.Keyring, error) {
	kb := keyring.NewInMemory(cdc)
	tmpPhrase := randPhrase(64)
	armored := cosmcrypto.EncryptArmorPrivKey(privKey, tmpPhrase, privKey.Type())
	err := kb.ImportPrivKey(name, armored, tmpPhrase)
	if err != nil {
		err = errors.Wrap(err, "failed to import privkey")
		return nil, err
	}

	return kb, nil
}

func KeyringForMnemonic(cdc codec.Codec, name, mnemonic string) (cosmtypes.AccAddress, keyring.Keyring, error) {
	kb := keyring.NewInMemory(cdc)
	info, err := kb.NewAccount(name, mnemonic, "", "m/44'/118'/0'/0/0", hd.Secp256k1)
	if err != nil {
		err = errors.Wrap(err, "failed to import privkey")
		return emptyCosmosAddress, nil, err
	}

	addr, err := info.GetAddress()
	if err != nil {
		return nil, nil, err
	}
	return addr, kb, nil
}

func randPhrase(size int) string {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	orPanic(err)

	return string(buf)
}

func orPanic(err error) {
	if err != nil {
		log.Panicln()
	}
}
