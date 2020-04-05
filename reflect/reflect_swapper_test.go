package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapper(t *testing.T) {
	a := []int64{1, 2, 3}
	fn := reflect.Swapper(a)

	fn(1, 2)
	assert.Equal(t, []int64{1, 3, 2}, a)
}
