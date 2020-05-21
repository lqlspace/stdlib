package client

import (
	"net/http"
	"net/url"

	"github.com/lqlspace/http/client/util"
)

type FormObj struct {

}

func (fo *FormObj) SendForm() (*http.Response, error) {
	urlPath := util.ADDRESS + `/form`
	form := url.Values{
		"ak": {"av1", "av2", "av3"},
	}
	form.Set("bk", "bv")

	return http.DefaultClient.PostForm(urlPath, form)
}
