package cryptox

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestMD5(t *testing.T) {
	dataBytes := []byte(`congratulations, girls!`)
	hash := md5.New()
	hash.Write(dataBytes)
	ciphertext := hex.EncodeToString(hash.Sum(nil))

	t.Logf("ciphertext = %s\n", ciphertext)
}
