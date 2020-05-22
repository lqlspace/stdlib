package client

import (
	"bytes"
	"net/http"

	"github.com/lqlspace/http/client/util"
)

type Json struct {}


func (js *Json) SendAndReceiveJson(data []byte) (*http.Response, error) {
	urlPath := util.ADDRESS + `/json`

	req, err :=  http.NewRequest("POST", urlPath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return  http.DefaultClient.Do(req)
}
