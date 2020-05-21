package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ParameterObj struct {

}

func (po *ParameterObj) AddRoutes() {
	http.HandleFunc("/str/get", po.StrGetWithQuery)
	http.HandleFunc("/str/post", po.StrPost)
}

func (po *ParameterObj) StrGetWithQuery(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	for key, val := range q {
		fmt.Printf("key = %s, val = %s\n", key, val)
	}

	fmt.Fprintf(w, "git params success: %s\n", r.URL.RawQuery)
}


func (po *ParameterObj) StrPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "can not parse body")
	}

	if string(body) == "welcome" {
		fmt.Fprintf(w, "yes,ok!")
	} else {
		fmt.Fprintf(w, "no, sorry!")
	}
}
