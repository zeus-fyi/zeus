package bls_signer

import blst "github.com/supranational/blst/bindings/go"

type PublicKey struct {
	p blst.P1Affine
}

func NewPubKey(pubKey blst.P1Affine) PublicKey {
	return PublicKey{pubKey}
}

// Serialize a public key into a LittleEndian byte slice.
func (k *PublicKey) Serialize() []byte {
	pubKeyBytes := k.p.Serialize()
	return pubKeyBytes
}

func PublicKeyFromBytes(b []byte) PublicKey {
	pubKey := new(blst.P1Affine).Deserialize(b)
	return PublicKey{p: *pubKey}
}
