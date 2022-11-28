package bls_signer

import blst "github.com/supranational/blst/bindings/go"

type SecretKey struct {
	s *blst.SecretKey
}

func NewSecretKey(sk *blst.SecretKey) SecretKey {
	return SecretKey{sk}
}

// Serialize a secret key into a LittleEndian byte slice.
func (k *SecretKey) Serialize() []byte {
	privKeyBytes := k.s.Serialize()
	return privKeyBytes
}

func SecretKeyFromBytes(b []byte) SecretKey {
	secKey := new(blst.SecretKey).Deserialize(b)
	return SecretKey{s: secKey}
}
