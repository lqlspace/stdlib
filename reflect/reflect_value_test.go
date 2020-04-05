package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueIsCanAddr(t *testing.T) {
	//int64
	var ai int64 = 3
	aipv := reflect.ValueOf(&ai)

	aiv := reflect.Indirect(aipv)

	if aiv.CanAddr() {
		t.Logf("can address")
		if aiv.CanSet() {
			aiv.Set(reflect.ValueOf(int64(20)))
			t.Logf("ai = %d\n", aiv.Int())
		}
	} else {
		t.Logf("can not address")
	}
}

//IsNil函数的入参只能是具有指针类型特点的类型
func TestValueIsNil(t *testing.T) {
	var isTrue bool

	//1.1 pointer,且不为nil
	a := new(Employee)
	av := reflect.ValueOf(a)
	isTrue = av.IsNil()
	assert.False(t, isTrue)

	//1.2 pointer,且为nil指针
	var b *Employee
	bv := reflect.ValueOf(b)
	isTrue = bv.IsNil()
	assert.True(t, isTrue)

	//2.1 slice，且为nil指针
	var as []int
	asv := reflect.ValueOf(as)
	isTrue = asv.IsNil()
	assert.True(t, isTrue)

	//2.2 slice,且为非Nil指针
	as = make([]int, 0)
	asv = reflect.ValueOf(as)
	isTrue = asv.IsNil()
	assert.False(t, isTrue)

	//3.1 func，且为nil
	var fn func() error
	value := reflect.ValueOf(fn)
	isTrue = value.IsNil()
	assert.True(t, isTrue)

	//3.2 func,非nil
	fn = func() error {
		return nil
	}

	value = reflect.ValueOf(fn)
	isTrue = value.IsNil()
	assert.False(t, isTrue)

}

//isValid是对reflect.Value本身的判定，不是对它wrap的值的判定
func TestIsValid(t *testing.T) {
	var isTrue bool

	var a int
	val := reflect.ValueOf(&a)
	isTrue = val.IsValid()
	assert.True(t, isTrue)

	//未init的pointer本身获得的val不为Nil,但indirect后指向空
	var b *int
	val = reflect.ValueOf(b)
	isTrue = val.IsValid()
	assert.True(t, isTrue)
	val1 := reflect.Indirect(val)
	isTrue = val1.IsValid()
	assert.False(t, isTrue)

	//此处reflect.Value取0值
	var c reflect.Value
	isTrue = c.IsValid()
	assert.False(t, isTrue)

	d := reflect.ValueOf(nil)
	isTrue = d.IsValid()
	assert.False(t, isTrue)
}

func TestIsZero(t *testing.T) {
	var isTrue bool

	var a int
	value := reflect.ValueOf(a)
	isTrue = value.IsZero()
	assert.True(t, isTrue)

	pa := reflect.ValueOf(&a)
	pav := reflect.Indirect(pa)
	if pav.CanAddr() {
		if pav.CanSet() {
			pav.Set(reflect.ValueOf(20))
		}
	}

	value = reflect.ValueOf(a)
	isTrue = value.IsZero()
	assert.False(t, isTrue)
}

func TestIsValidIsNilIsZero(t *testing.T) {
	var a, c *int
	var b int

	//b = 20
	c = &b

	av := reflect.ValueOf(a)
	cv := reflect.ValueOf(c)

	if av.IsValid() {
		if av.IsNil() {
			t.Logf("av is nil")
		} else {
			t.Log("av is not nil")
		}
	}

	if cv.IsValid() {
		if cv.IsNil() {
			t.Log("cv is nil")
		} else {
			t.Log("cv is not nil")
		}
	}

	avi := reflect.Indirect(av)
	cvi := reflect.Indirect(cv)

	if avi.IsValid() {
		t.Log("avi is valid")
		if avi.CanAddr() {
			t.Log("avi can addr")
			if avi.IsNil() {
				t.Log("avi is nil")
			} else {
				t.Log("avi is not nil")
			}
		} else {
			t.Log("avi can not addr")
		}
	} else {
		t.Log("avi is not valid")
	}

	if cvi.IsValid() {
		t.Log("cvi is valid")
		if cvi.CanAddr() {
			t.Log("cvi can addr")
			if cvi.IsZero() {
				t.Log("cvi is zero")
			} else {
				t.Log("cvi is not zero")
			}
		} else {
			t.Log("cvi can not addr")
			if cvi.IsZero() {
				t.Log("cvi is zero")
			} else {
				t.Log("cvi is not zero")
			}
		}
	} else {
		t.Log("cvi is not valid")
	}

}
