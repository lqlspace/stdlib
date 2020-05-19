package server

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

const (
	ADDR = `:8080`
)

func TestServerSimple_ServeSimple(t *testing.T) {
	ss := new(SimpleConn)

	ss.AddRoutes()

	fmt.Printf("Starting SimpleConnect server at  : %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, nil))
}
