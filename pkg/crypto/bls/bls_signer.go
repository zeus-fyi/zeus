package bls_signer

import blst "github.com/supranational/blst/bindings/go"

type Signature = blst.P2Affine
type AggregateSignature = blst.P2Aggregate
type AggregatePublicKey = blst.P1Aggregate

// alternate uses var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")
var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

func (k *KeyBLS) Sign(msg []byte) *Signature {
	sig := new(Signature).Sign(k.SecretKey.s, msg, dst)
	return sig
}

func (k *KeyBLS) Verify(sig Signature, msg []byte) bool {
	if !sig.Verify(true, &k.PublicKey.p, true, msg, dst) {
		return false
	} else {
		return true
	}
}
