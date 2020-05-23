package client

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lqlspace/http/client/util"
)

func TestFormObj_SendForm(t *testing.T) {
	fo := new(FormObj)

	urlPath := util.ADDRESS + `/form`
	rsp, err := fo.SendForm(urlPath)
	assert.Nil(t, err)
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)

	t.Logf("body = %s\n", string(body))
}


func TestFormObj_UploadFile(t *testing.T) {
	fo := new(FormObj)

	urlPath := util.ADDRESS + `/upload/file`
	rsp, err := fo.UploadFile("doc/jackie.txt", urlPath)
	assert.Nil(t, err)

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)

	// 传递的纯字符串，直接打印
	t.Logf("%s\n", string(body))
}
