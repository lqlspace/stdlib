package server

import (
	"fmt"
	"net/http"
)

type FormObj struct {

}


func (fo *FormObj) AddRoutes() {
	http.HandleFunc("/form", fo.HandleForm)
}


func (fo *FormObj) HandleForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, val := range r.Form {
		fmt.Printf("%s: %s\n", key, val)
		fmt.Fprintf(w, "%s: %s\n", key, val)
	}
}
