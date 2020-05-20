package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleConn_GetUr(t *testing.T) {
	sc := new(SimpleConn)
	rspInfo, err := sc.GetUrl()
	assert.Nil(t, err)

	t.Log(rspInfo)
}

func TestSimpleConn_GetHeader(t *testing.T) {
	sc := new(SimpleConn)

	rsp, err := sc.GetHeader()
	assert.Nil(t, err)
	assert.NotNil(t, rsp, rsp.Body)

	defer rsp.Body.Close()

	t.Logf("rsp = %#v", rsp)


}

//http.Head方法，http.Body为http.noBody{}
func TestSimpleConn_GetHeaderByHead(t *testing.T) {
	sc := new(SimpleConn)

	rsp, err := sc.GetHeaderByHead()
	assert.Nil(t, err)
	assert.Nil(t, rsp.Body)

	t.Logf("rsp = %#v", rsp)
}

func TestSimpleConn_GetMethod(t *testing.T) {
	sc := new(SimpleConn)

	data, err :=  sc.GetMethod()
	assert.Nil(t, err)
	t.Logf("data: %s\n", data)
}
