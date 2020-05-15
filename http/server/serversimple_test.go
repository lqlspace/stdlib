package server

import (
	"log"
	"net/http"
	"testing"

	"xiao/common/logx"
)

const (
	ADDR = `:8080`
)

func TestServerSimple_ServeSimple(t *testing.T) {
	ss := new(ServerSimple)

	ss.ServeSimple()

	logx.Infof("Starting SimpleConnect server at  : %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, nil))
}
