package aegis_random

import (
	"crypto/rand"

	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
)

func GenerateMnemonic() (string, error) {
	entropy := make([]byte, 32)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate entropy for wallet mnemonic")
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate wallet mnemonic")
	}
	return mnemonic, err
}
