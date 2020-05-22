package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Json struct {

}


func (js *Json) AddRoutes() {
	http.HandleFunc("/json", js.HandleJsonData)
}


func (js *Json) HandleJsonData(w http.ResponseWriter, r *http.Request)  {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err !=  nil {
		fmt.Fprintf(w, "parse json data failed: %s", err)
		return
	}

	for k, v := range r.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	body["response"] = "have received!"

	json.NewEncoder(w).Encode(body)
}
