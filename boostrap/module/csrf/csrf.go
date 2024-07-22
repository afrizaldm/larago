package csrf

import (
	"crypto/rand"
	"encoding/base64"
)

// Fungsi untuk menghasilkan token CSRF
func GenerateCSRFToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}
