package client

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameterObj_GetWithParams(t *testing.T) {
	po := new(ParameterObj)
	rsp, err := po.StrGetWithQuery()
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)

	t.Logf("body = %s\n", string(body))
}


func TestParameterObj_StrPost(t *testing.T) {
	po := new(ParameterObj)
	rsp, err := po.StrPost()
	assert.Nil(t, err)

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)

	t.Logf("body = %s\n", string(body))
}
