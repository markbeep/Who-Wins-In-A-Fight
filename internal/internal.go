package internal

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

// Generate a random token of a given length.
// len is the amount of random bytes to generate.
func GenerateToken(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed generating a random token. err = %s", err)
	}
	return strings.ReplaceAll(base64.URLEncoding.EncodeToString(b), "=", ""), nil
}
