package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type ecb struct {
	block cipher.Block
	blockSize int
}

func newECB(block cipher.Block) *ecb {
	return &ecb{
		block: block,
		blockSize: block.BlockSize(),
	}
}

type ecbEncrypter ecb


func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (e *ecbEncrypter) BlockSize() int {
	return e.blockSize
}

func (e *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		fmt.Printf("crypto: input not full blocks")
		return
	}

	if len(dst) < len(src) {
		fmt.Printf("crypto: output smaller than input")
	}

	for len(src) > 0 {
		e.block.Encrypt(dst, src[:e.blockSize])
		src  = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}


type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (e *ecbDecrypter) BlockSize() int {
	return e.blockSize
}

func (e *ecbDecrypter ) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		fmt.Printf("crypto: input not full blocks")
		return
	}

	if len(dst) < len(src) {
		fmt.Printf("crypto: output smaller than input")
		return
	}

	for len(src) >  0 {
		e.block.Decrypt(dst, src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}


func EcbEncrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("encrypt key error")
	}

	paddedText, err := pkcs7Padding(plaintext, block.BlockSize())
	if err != nil {
		return nil, fmt.Errorf("padding error: %s\n", err)
	}

	ciphertext := make([]byte, len(paddedText))
	encrypter := NewECBEncrypter(block)
	encrypter.CryptBlocks(ciphertext, paddedText)

	return ciphertext, nil
}



func EcbDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("invalid key:%x\n", key)
	}

	decrypter := NewECBDecrypter(block)
	plaintext := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(plaintext, ciphertext)

	return pkcs7Unpadding(plaintext, block.BlockSize())
}
