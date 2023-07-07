package aegis_random

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomPassword(length int) (string, error) {
	// Define the character set that the password can be composed of
	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+-=[]{}|;:,.<>?`~"

	// Convert the character set to a slice of runes for easier random selection
	charSetRunes := []rune(charSet)

	// Use a cryptographically secure random number generator to select characters from the character set
	// and build up the password string
	password := make([]rune, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSetRunes))))
		if err != nil {
			return "", err
		}
		password[i] = charSetRunes[randomIndex.Int64()]
	}

	return string(password), nil
}
