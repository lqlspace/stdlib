package cryptox

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESECB(t *testing.T) {
	key := sha256.Sum256([]byte(`key`))

	plaintext := []byte(`jackie`)

	ciphertext, err := EcbEncrypt(key[:], plaintext)
	assert.Nil(t, err)

	newText, err := EcbDecrypt(key[:], ciphertext)
	assert.Nil(t, err)

	assert.Equal(t, newText, plaintext)
}

