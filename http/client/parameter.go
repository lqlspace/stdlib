package client

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/lqlspace/http/client/util"
)

type ParameterObj struct {

}


func (po *ParameterObj) StrGetWithQuery() (*http.Response, error) {
	urlPath := util.ADDRESS + `/str/get`
	req, err :=  http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add("name", "allen")
	query.Add("password", "123")
	req.URL.RawQuery = query.Encode()

	return http.DefaultClient.Do(req)
}


func (po *ParameterObj) StrGetWithPath() (*http.Response, error) {
	//todo
	return nil, nil
}


func (po *ParameterObj) StrPost() (*http.Response, error) {
	urlPath := util.ADDRESS + `/str/post`
	req, err := http.NewRequest("POST", urlPath, strings.NewReader("welcome"))
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
