package client

import (
	"io/ioutil"
	"net/http"
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


func TestFormObj_UploadJpg(t *testing.T) {
	fo := new(FormObj)

	urlPath := util.ADDRESS + `/save/jpg`
	rsp, err := fo.UploadJpg("doc/image.jpg", urlPath)
	assert.Nil(t, err)

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		t.Errorf("response error")
	} else {
		body, err := ioutil.ReadAll(rsp.Body)
		assert.Nil(t, err)

		t.Logf("%s\n", string(body))
	}
}


func TestFormObj_UploadVideo(t *testing.T) {
	fo := new(FormObj)

	urlPath := util.ADDRESS + `/save/video`
	rsp, err := fo.UploadVideo("doc/sunrise.mp4", urlPath)
	assert.Nil(t, err)

	defer rsp.Body.Close()
	if rsp.StatusCode !=  http.StatusOK {
		t.Errorf("upload video failed: %s\n", err)
	} else {
		body, err := ioutil.ReadAll(rsp.Body)
		assert.Nil(t, err)

		t.Logf("%s\n", string(body))
	}
}
