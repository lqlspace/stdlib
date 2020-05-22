package server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestJson_HandleJsonData(t *testing.T) {
	jsObj := new(Json)

	jsObj.AddRoutes()

	fmt.Printf("server start at %s ...\n", ADDR)
	http.ListenAndServe(ADDR, nil)
}
