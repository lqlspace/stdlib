package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES(t *testing.T) {
	text := []byte("My super secret code stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	assert.Nil(t, err)

	gcm, err := cipher.NewGCM(c)
	assert.Nil(t, err)

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	assert.Nil(t, err)

	ciphertext := hex.EncodeToString(gcm.Seal(nonce, nonce, text, nil))
	t.Log(ciphertext)
}


func TestAESCBC(t *testing.T) {
	//1. 随机生成256bit key，更加key生成block
	key := sha256.Sum256([]byte("jackie"))
	block, err := aes.NewCipher(key[:])
	assert.Nil(t, err)

	//2. plaintext，长度是blocksize(16)的整数倍
	text := bytes.Repeat([]byte("a"), 96)
	ciphertext := make([]byte, len(text))

	//3. initialization vector，作料
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	assert.Nil(t, err)

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext, text)

	t.Logf("ciphertext: %x\n", ciphertext)

	newText := make([]byte, len(ciphertext))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(newText, ciphertext)

	if  bytes.Equal(newText, text) {
		t.Log("encrypt decrypt process is ok")
	}
	assert.Equal(t, newText, text)
}



func TestAesCbc(t *testing.T) {
	//1. 生成Key 128位
	key256 := sha256.Sum256([]byte(`jackie`))

	block, err := aes.NewCipher(key256[:])
	assert.Nil(t, err)

	//2. 生成iv
	iv := make([]byte, aes.BlockSize)
	n, err := rand.Read(iv)
	assert.Nil(t, err)
	assert.Equal(t, 16, n)

	text := []byte(`welcome`)

	textPadding, err  := pkcs7Padding(text, aes.BlockSize)
	assert.Nil(t, err)

	cipherText := make([]byte, len(textPadding))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, textPadding)

	newText := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(newText, cipherText)
	newText, err = pkcs7Unpadding(newText, aes.BlockSize)
	assert.Nil(t, err)
	assert.Equal(t, text, newText)
}


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

