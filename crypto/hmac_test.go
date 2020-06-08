package cryptox

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestHMAC(t *testing.T) {
	secret := "mysecret"
	data :=  "data"

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	ciphertext := hex.EncodeToString(h.Sum(nil))

	t.Logf("len: %d, result : %s\n", len(ciphertext), ciphertext)
}
