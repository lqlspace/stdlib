package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESGCM(t *testing.T) {
	plaintext := []byte(`jakie`)

	// 96bits
	nonce := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, nonce)
	assert.Nil(t, err)

	ciphertext, err := encryptAESGCM(plaintext, nonce)
	assert.Nil(t, err)

	newText, err := decryptAESGCM(ciphertext, nonce)
	assert.Nil(t, err)

	assert.Equal(t, plaintext, newText)
}

func encryptAESGCM(plaintext, nonce []byte) ([]byte, error) {
	//256bits
	key, err := hex.DecodeString(`6368616e676520746869732070617373776f726420746f206120736563726574`)
	if err != nil {
		return nil, fmt.Errorf("retrieve key error: %s\n", err)
	}

	//128bits block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("invalid key: %s\n", key)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("invalid block: %s\n", err)
	}



	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nil
}


func decryptAESGCM(ciphertext, nonce []byte) ([]byte, error) {
	key, err := hex.DecodeString(`6368616e676520746869732070617373776f726420746f206120736563726574`)
	if err != nil {
		return nil, fmt.Errorf("retrieve key  failed: %s\n", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("invalid key: %s\n", key)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("invalid block: %s\n", err)
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt process:%s\n", err)
	}

	return plaintext, nil
}
