package client

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson_SendAndReceiveJson(t *testing.T) {
	data := map[string]interface{}{
		"name": "Jackie",
		"age": 20,
		"nested": []string{
			"yes, ok",
			"no, sorry",
		},
	}

	dBytes, err  := json.Marshal(data)
	assert.Nil(t, err)

	jsonObj := new(Json)
	rsp, err := jsonObj.SendAndReceiveJson(dBytes)
	assert.Nil(t,  err)
	defer rsp.Body.Close()

	var rspData map[string]interface{}
	err  = json.NewDecoder(rsp.Body).Decode(&rspData)
	assert.Nil(t, err)

	for key, val := range rspData {
		t.Logf("%s: %v\n", key, val)
	}
}
