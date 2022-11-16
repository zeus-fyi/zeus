package aegis_random

import (
	"crypto/rand"
	"encoding/hex"
)

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// Hex generates a random hex string with length of n
// e.g: 67aab2d956bd7cc621af22cfb169cba8
func Hex(n int) string {
	return hex.EncodeToString(Bytes(n))
}
