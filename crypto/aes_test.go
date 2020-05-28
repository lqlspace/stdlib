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

	textPadding  := pkcs5Padding(text, aes.BlockSize)

	cipherText := make([]byte, len(textPadding))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, textPadding)

	newText := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(newText, cipherText)
	newText, err = pkcs5Unpadding(newText, aes.BlockSize)
	assert.Nil(t, err)
	assert.Equal(t, text, newText)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Unpadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding >= length || unpadding > blockSize {
		return nil, errors.New("there's error")
	}

	return src[:length-unpadding], nil
}
