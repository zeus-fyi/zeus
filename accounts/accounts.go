package accounts

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

type Account struct {
	key *ecdsa.PrivateKey
}

func CreateAccount() (*Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Err(err).Msg("CreateAccount: crypto.GenerateKey()")
		return nil, err
	}
	return &Account{
		key: key,
	}, nil
}

func ParsePrivateKey(pkHex string) (*Account, error) {
	fromPK := strings.TrimPrefix(pkHex, "0x")
	key, err := crypto.HexToECDSA(fromPK)
	if err != nil {
		log.Err(err).Msg("ParsePrivateKey: HexToECDSA")
		return nil, err
	}
	return &Account{
		key: key,
	}, nil
}

func (a *Account) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(data, a.key)
}

func (a *Account) VerifySignature(pubkey Address, data, sig []byte) (bool, error) {
	pubkeyVal, err := crypto.SigToPub(data[:], sig[:])
	if err != nil {
		log.Err(err).Msg("VerifySignature: SigToPub")
		return false, err
	}
	addr := crypto.PubkeyToAddress(*pubkeyVal)
	return addr.String() == pubkey.String(), nil
}

func (a *Account) EcdsaPrivateKey() *ecdsa.PrivateKey {
	return a.key
}

func (a *Account) Address() Address {
	return Address(crypto.PubkeyToAddress(a.key.PublicKey))
}

func (a *Account) PublicKey() string {
	return crypto.PubkeyToAddress(a.key.PublicKey).Hex()
}

func (a *Account) PrivateKey() string {
	return "0x" + hex.EncodeToString(crypto.FromECDSA(a.key))
}

func (a *Account) EcdsaPublicKey() *ecdsa.PublicKey {
	privateKey := a.EcdsaPrivateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err := errors.New("error casting public key to ECDSA")
		log.Panic().Err(err).Msg("EcdsaPublicKey")
		panic(err)
	}
	return publicKeyECDSA
}
