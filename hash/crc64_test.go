package hashx

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc64"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestCRC64(t *testing.T) {
	org := []byte(`welcome, guys!`)
	table := crc64.MakeTable(crc64.ISO)

	hash64 := crc64.New(table)
	_, err := hash64.Write(org)
	assert.Nil(t, err)

	cs := hash64.Sum64()

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, cs)


	t.Logf("size: %d, content: %s\n", unsafe.Sizeof(cs), hex.EncodeToString(buf))

}
