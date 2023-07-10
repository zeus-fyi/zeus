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
	key            *ecdsa.PrivateKey
	nonceIncrement uint64
}

func CreateAccount() (*Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Err(err).Msg("CreateAccount: crypto.GenerateKey()")
		return nil, err
	}
	return &Account{
		key:            key,
		nonceIncrement: 0,
	}, nil
}

func CreateAccountFromPkey(pkey *ecdsa.PrivateKey) (*Account, error) {
	return &Account{
		key: pkey,
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
	signedData, err := crypto.Sign(data, a.key)
	if err != nil {
		log.Err(err).Msg("Sign: crypto.Sign")
		return nil, err
	}
	/*
		// EcRecover returns the address for the account that was used to create the signature.
		// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
		// the address of:
		// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
		// addr = ecrecover(hash, signature)
		//
		// Note, the signature must conform to the secp256k1 curve R, S and V values, where
		// the V value must be be 27 or 28 for legacy reasons.
	*/
	signedData[64] += 27
	return signedData, nil
}

func (a *Account) VerifySignature(pubkey Address, data, sig []byte) (bool, error) {
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
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

func (a *Account) GetNonceOffset() uint64 {
	return a.nonceIncrement
}

func (a *Account) IncrementLocalNonce() {
	a.nonceIncrement++
}

func (a *Account) ResetLocalNonce() {
	a.nonceIncrement = 0
}
