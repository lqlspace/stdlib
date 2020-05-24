package hashx

import (
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRC32(t *testing.T) {
	dataBytes := []byte(`hello, everyone!`)
	cs1 := crc32.ChecksumIEEE(dataBytes)

	cs2 := crc32.Checksum(dataBytes, crc32.IEEETable)

	ieeeTable := crc32.MakeTable(crc32.IEEE)
	cs3 := crc32.Checksum(dataBytes, ieeeTable)

	//step1: create hash.Hash object
	hash := crc32.New(ieeeTable)
	//step2: get original data
	hash.Write(dataBytes)
	//step3: generate and return crc32 checksum
	cs4 := hash.Sum32()

	assert.EqualValues(t, cs1, cs2)
	assert.EqualValues(t, cs1, cs3)
	assert.EqualValues(t, cs3, cs4)
}
