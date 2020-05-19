package client

import (
	"io/ioutil"
	"net/http"

	"github.com/lqlspace/http/client/util"
)

type SimpleConn struct {
}

func (cs *SimpleConn) GetUrl() (string, error) {
	urlPath := util.ADDRESS + `/url`
	rsp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


func (cs *SimpleConn) GetHeader()  (string, error) {
	urlPath := util.ADDRESS + `/header`

	req, err := http.NewRequest("Get", urlPath, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "abcdefg")
	req.Header.Set("name", "allen")

	rsp, err := http.DefaultClient.Do(req)
	if err !=  nil {
		return "", err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err !=  nil {
		return "", err
	}

	return string(body), err
}


func (cs *SimpleConn) GetMethod() (string, error) {
	urlPath := util.ADDRESS + `/method`

	rsp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
