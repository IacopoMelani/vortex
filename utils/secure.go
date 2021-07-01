package utils

import "crypto/rand"

const rune = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// SecureRandomString - Returns a secure random string
func SecureRandomString(length uint) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = rune[b%byte(len(rune))]
	}

	return string(bytes), nil
}
