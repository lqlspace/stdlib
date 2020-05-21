package server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFormObj_HandleForm(t *testing.T) {
	fo := new(FormObj)

	fo.AddRoutes()

	fmt.Printf("server start at %s ...\n", ADDR)
	http.ListenAndServe(ADDR, nil)
}
