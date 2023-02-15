package age_encryption

import (
	"filippo.io/age"
)

type Age struct {
	agePrivateKey string
	agePublicKey  string
}

func NewAge(privKey, pubKey string) Age {
	a := Age{
		agePrivateKey: privKey,
		agePublicKey:  pubKey,
	}
	return a
}

// GenerateNewKeyPair returns publicKey, privateKey (order of returns)
func GenerateNewKeyPair() (string, string) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		panic(err)
	}

	pubkey := identity.Recipient().String()
	privKey := identity.String()
	return pubkey, privKey
}

func GenerateNewAgeCredentials() Age {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return Age{}
	}

	pubkey := identity.Recipient().String()
	privKey := identity.String()
	return Age{
		agePrivateKey: privKey,
		agePublicKey:  pubkey,
	}
}
