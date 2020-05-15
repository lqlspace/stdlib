package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ClientSimple struct {
}

func (cs *ClientSimple) ConnGet() (string, error) {
	rsp, err := http.Get("http://127.0.0.1:8080/user")
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

func (cs *ClientSimple) ConnPost() (string, error) {
	resp, err := http.Post("http://127.0.0.1:8080/user", "", nil)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cs *ClientSimple) ConnPostWithParams() (string, error) {
	//Tips：使用该形式，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。
	rsp, err := http.Post("http://127.0.0.1:8080/user",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cs *ClientSimple) ConnPostForm() (string, error) {
	rsp, err := http.PostForm("http://127.0.0.1:8080/user", url.Values{"key": {"Value"}, "id": {"123"}})
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
