package cryptox

import (
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPEMCreate(t *testing.T) {
	file, err := os.Create("test.pem")
	assert.Nil(t, err)

	data := []byte("pem tst")
	t.Logf("len = %d, data = %x, datab = %b\n", len(data), data, data)

	t.Logf("%d, %d\n", 0b00000111, 0b00000011)

	err = pem.Encode(file, &pem.Block{
		Type:    "JACKIE CHAN",
		Headers: map[string]string{
			"aaa": "vaaa",
			"bbb": "vbbb",
		},
		Bytes:   data,
	})

	assert.Nil(t, err)
}
