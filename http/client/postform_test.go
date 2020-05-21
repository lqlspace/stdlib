package client

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormObj_SendForm(t *testing.T) {
	fo := new(FormObj)
	rsp, err := fo.SendForm()
	assert.Nil(t, err)
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)

	t.Logf("body = %s\n", string(body))


}
