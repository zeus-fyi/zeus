package bls_signer

import (
	"crypto/rand"

	blst "github.com/supranational/blst/bindings/go"
)

type KeyBLS struct {
	PublicKey blst.P1Affine
	SecretKey *blst.SecretKey
}

func NewKeyBLS() KeyBLS {
	var ikm [32]byte
	_, err := rand.Read(ikm[:])
	if err != nil {
		panic(err)
	}
	sk := blst.KeyGen(ikm[:])
	pk := new(PublicKey).From(sk)
	k := KeyBLS{PublicKey: *pk, SecretKey: sk}
	return k
}
