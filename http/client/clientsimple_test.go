package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientSimple_ConnGet(t *testing.T) {
	cs := new(ClientSimple)
	rspInfo, err := cs.ConnGet()
	assert.Nil(t, err)

	t.Log(rspInfo)
}

func TestClientSimple_ConnPost(t *testing.T) {
	cs := new(ClientSimple)
	rspInfo, err := cs.ConnPost()
	assert.Nil(t, err)

	t.Log(rspInfo)
}

func TestClientSimple_ConnPostWithParams(t *testing.T) {
	cs := new(ClientSimple)
	rspInfo, err := cs.ConnPost()
	assert.Nil(t, err)

	t.Log(rspInfo)
}

func TestClientSimple_ConnPostForm(t *testing.T) {
	cs := new(ClientSimple)

	rsp, err := cs.ConnPostForm()
	assert.Nil(t, err)

	t.Log(rsp)
}
