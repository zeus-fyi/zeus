package bls_signer

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	blst "github.com/supranational/blst/bindings/go"
)

type Account struct {
	KeyBLS
}

type KeyBLS struct {
	PublicKey PublicKey
	SecretKey SecretKey
}

func NewSignerBLSFromExistingKey(skStr string) Account {
	var ba [32]byte
	skByteArr, err := hex.DecodeString(strings.TrimPrefix(skStr, "0x"))
	if err != nil {
		panic(err)
	}
	copy(ba[:], skByteArr)
	sk := new(blst.SecretKey).Deserialize(ba[:])
	pk := new(blst.P1Affine).From(sk)
	k := KeyBLS{PublicKey: NewPubKey(*pk), SecretKey: NewSecretKey(sk)}
	return Account{k}
}

func NewSignerBLS() Account {
	return Account{NewKeyBLS()}
}

func NewKeyBLS() KeyBLS {
	var ikm [32]byte
	_, err := rand.Read(ikm[:])
	if err != nil {
		panic(err)
	}
	sk := blst.KeyGen(ikm[:])
	pk := new(blst.P1Affine).From(sk)
	k := KeyBLS{PublicKey: NewPubKey(*pk), SecretKey: NewSecretKey(sk)}
	return k
}
