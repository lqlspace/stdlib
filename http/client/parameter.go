package client

import (
	"net/http"
	"net/url"

	"github.com/lqlspace/http/client/util"
)

type ParameterObj struct {

}


func (po *ParameterObj) GetWithQuery() (*http.Response, error) {
	urlPath := util.ADDRESS + `/params`
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


func (po *ParameterObj) GetWithPath() (*http.Response, error) {
	//todo
	return nil, nil
}
