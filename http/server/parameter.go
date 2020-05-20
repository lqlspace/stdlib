package server

import (
	"fmt"
	"net/http"
)

type ParameterObj struct {

}

func (po *ParameterObj) AddRoutes() {
	http.HandleFunc("/params", po.ConnWithParams)
}

func (po *ParameterObj) ConnWithParams(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	for key, val := range q {
		fmt.Printf("key = %s, val = %s\n", key, val)
	}

	fmt.Fprintf(w, "git params success: %s\n", r.URL.RawQuery)
}
