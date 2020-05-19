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

	data, err := sc.GetHeader()
	assert.Nil(t, err)

	t.Log(data)
}


func TestSimpleConn_GetMethod(t *testing.T) {
	sc := new(SimpleConn)

	data, err :=  sc.GetMethod()
	assert.Nil(t, err)
	t.Logf("data: %s\n", data)
}
