package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayEqual(t *testing.T) {
	// 两个均为nil值(貌似没有？)

	// 其中一个为nil值（貌似没有？）

	// 类型不同，注：golang中数组个数算作数组类型的一部分
	var an [3]int64
	var bnn [2]int64
	equal := reflect.DeepEqual(an, bnn)
	assert.False(t, equal)

	// 类型相同，个数相等，但均为0值
	var bn [3]int64
	equal = reflect.DeepEqual(an, bn)
	assert.True(t, equal)

	var cn [0]int64
	var dn [0]int64
	equal = reflect.DeepEqual(cn, dn)
	assert.True(t, equal)

	// 个数相等，但元素不同
	as := [3]string{"1", "2"}
	bs := [3]string{"1", "3", "4"}
	equal = reflect.DeepEqual(as, bs)
	assert.False(t, equal)

	// 个数相等，值相同
	as = [3]string{"1", "2", "3"}
	bs = [3]string{"1", "2", "3"}
	equal = reflect.DeepEqual(as, bs)
	assert.True(t, equal)
}

func TestSliceEqual(t *testing.T) {
	var equal bool

	// 均为Nil值(此处有相同底层数组)
	var an []int64
	var bn []int64
	equal = reflect.DeepEqual(an, bn)
	assert.True(t, equal)

	// 一个为空
	cn := make([]int64, 0)
	equal = reflect.DeepEqual(an, cn)
	assert.False(t, equal)

	// 均非nil，但长度不同
	s := []int64{1, 2, 3, 4, 5}
	as := s[1:2]
	bs := s[1:4]
	equal = reflect.DeepEqual(as, bs)
	assert.False(t, equal)

	// 均非nil，长度相同，但值不同
	cs := s[2:3]
	ds := s[1:2]
	equal = reflect.DeepEqual(cs, ds)
	assert.False(t, equal)

	// 均非nil，长度相同，值相同，但不是相同的底层数组
	x := []int64{2, 3, 4, 6, 2, 3}
	es := x[1:2]
	equal = reflect.DeepEqual(cs, es)
	assert.True(t, equal)

	// 均非nil，长度相同，值相同，相同的底层数组
	ls := x[:2]
	ks := x[4:]
	equal = reflect.DeepEqual(ls, ks)
	assert.True(t, equal)
}

type Student struct {
	Name    string
	Age     int
	Profile *Profile
}

type Profile struct {
	High int
	Male bool
}

func TestStructEqual(t *testing.T) {
	var equal bool

	// 均为0值 (递归比较结构体中每个字段)
	var a Student
	var b Student
	equal = reflect.DeepEqual(a, b)
	assert.True(t, equal)

	// 其中一个内部字段初始化
	var c Student
	c.Profile = new(Profile)
	equal = reflect.DeepEqual(a, c)
	assert.False(t, equal)
}

func TestPtrEqual(t *testing.T) {
	var equal bool

	//相同类型0值的指针相同
	var a int64
	var b int64
	pa := &a
	pb := &b
	equal = reflect.DeepEqual(pa, pb)
	assert.True(t, equal)

	//相同类型，其中一个非0值
	var c int64 = 3
	pc := &c
	equal = reflect.DeepEqual(pa, pc)
	assert.False(t, equal)

	//不同类型，均为0值
	var d string
	pd := &d
	equal = reflect.DeepEqual(pa, pd)
	assert.False(t, equal)
}

func TestMapEqual(t *testing.T) {
	var equal bool

	//均取nil值
	var a map[int64]int64
	var b map[int64]int64
	equal = reflect.DeepEqual(a, b)
	assert.True(t, equal)

	//一个nil，一个0值
	c := make(map[int64]int64)
	equal = reflect.DeepEqual(a, c)
	assert.False(t, equal)

	//两个0值
	d := make(map[int64]int64)
	equal = reflect.DeepEqual(c, d)
	assert.True(t, equal)

	//两个不同类型0值
	e := make(map[int64]int)
	equal = reflect.DeepEqual(d, e)
	assert.False(t, equal)

	d[1] = 2
	equal = reflect.DeepEqual(c, d)
	assert.False(t, equal)
}
