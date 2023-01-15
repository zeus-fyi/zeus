package age_encryption

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
