package server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParameterObj_ConnWithParams(t *testing.T) {
	po := new(ParameterObj)
	po.AddRoutes()

	fmt.Printf("conn with  params startup: %s...\n", ADDR)
	http.ListenAndServe(ADDR, nil)
}
