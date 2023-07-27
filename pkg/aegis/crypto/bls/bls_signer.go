package bls_signer

import blst "github.com/supranational/blst/bindings/go"

type Signature struct {
	s *blst.P2Affine
}
type AggregateSignature = blst.P2Aggregate
type AggregatePublicKey = blst.P1Aggregate

// alternate uses var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")
var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

func (k *KeyBLS) Sign(msg []byte) *Signature {
	sig := new(blst.P2Affine).Sign(k.SecretKey.s, msg, dst)
	return &Signature{s: sig}
}

func (k *KeyBLS) Verify(sig Signature, msg []byte) bool {
	if !sig.s.Verify(true, &k.PublicKey.p, true, msg, dst) {
		return false
	} else {
		return true
	}
}

func (s *Signature) Serialize() []byte {
	return s.s.Serialize()
}

func (s *Signature) String() string {
	return ConvertBytesToString(s.s.Serialize())
}
