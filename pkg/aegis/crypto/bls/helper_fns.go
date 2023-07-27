package bls_signer

import "fmt"

func ConvertBytesToString(b []byte) string {
	return fmt.Sprintf("%x", b)
}
