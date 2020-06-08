package base64x

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeToken(t *testing.T) {
	token :=  `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEyNDk2NzksImlhdCI6MTU5MDA0MDA3OSwidXNyIjoiMTcyMCJ9.dpfTtNEFf5c2sQP2PkBzk81K_eQGI_Bzs6WzwxDLrek`

	segs := strings.Split(token, ".")

	firstSeg, err := DecodeSegment(segs[0])
	assert.Nil(t, err)

	t.Logf("header = %s\n", string(firstSeg))
	//var header map[string]interface{}
	//err = json.Unmarshal(firstSeg,&header)
	//assert.Nil(t, err)
	//
	//for key, val := range header {
	//	t.Logf("%s: %v\n", key, val)
	//}

	secondSeg, err := DecodeSegment(segs[1])
	assert.Nil(t, err)

	t.Logf("claims = %s\n", string(secondSeg))

}


func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
