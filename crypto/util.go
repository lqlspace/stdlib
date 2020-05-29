package cryptox

import (
	"bytes"
	"errors"
)

func pkcs5Padding(plaintext []byte) []byte {
	paddingLen := 8 - len(plaintext) % 8
	paddingText := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)

	return append(plaintext, paddingText...)
}

func pkcs5Unpadding(ciphertext []byte) ([]byte, error) {
	paddingLen := int(ciphertext[len(ciphertext)-1])

	if paddingLen < 0 || paddingLen >= 8 {
		return nil, errors.New("padding length should not larger than blockSize")
	}

	return ciphertext[:len(ciphertext)-paddingLen], nil
}


func pkcs7Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}

	if plaintext == nil || len(plaintext) == 0 {
		return nil, errors.New("invalid pkcs7 data")
	}

	n := blockSize - (len(plaintext)%blockSize)

	return append(plaintext, bytes.Repeat([]byte{byte(n)}, n)...), nil
}

// pkcs stands for "Public Key Cryptography Standards"
func pkcs7Unpadding(ciphertext []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}

	if ciphertext == nil || len(ciphertext) == 0 {
		return nil, errors.New("invalid pkcs7 data")
	}

	paddingLen := int(ciphertext[len(ciphertext)-1])
	if paddingLen < 0 || paddingLen >= blockSize {
		return nil, errors.New("invalid padding size")
	}

	return ciphertext[:len(ciphertext)-paddingLen], nil
}
