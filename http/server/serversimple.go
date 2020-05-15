package server

import (
	"fmt"
	"net/http"
)

type ServerSimple struct{}

func (ss *ServerSimple) ServeSimple() {
	http.HandleFunc("/user", ss.Handler)
}

func (ss *ServerSimple) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
