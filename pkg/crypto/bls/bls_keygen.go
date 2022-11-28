package bls_signer

import (
	"crypto/rand"

	blst "github.com/supranational/blst/bindings/go"
)

type Account struct {
	KeyBLS
}

type KeyBLS struct {
	PublicKey PublicKey
	SecretKey SecretKey
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
