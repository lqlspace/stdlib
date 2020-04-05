package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructTag(t *testing.T) {
	type Student struct {
		Name    string   `liu:"name" zhao:"victoriaName"`
		Age     int      `liu:"age" zhao:"victoriaAge"`
		Profile *Profile `liu:"profile" zhao:"victoriaProfile"`
	}

	type Profile struct {
		Gender int
	}

	var s Student
	st := reflect.TypeOf(s)
	//sv := reflect.ValueOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		name := field.Tag.Get(`liu`)
		age := field.Tag.Get(`zhao`)
		fmt.Println(field.Name, field.Type, name, age)
	}
}

func TestStructTagLookup(t *testing.T) {
	type Student struct {
		Name    string   `liu:"name" zhao:"victoriaName"`
		Age     int      `liu:"-" zhao:"victoriaAge"`
		Profile *Profile `liu:"profile" zhao:"victoriaProfile"`
	}

	type Profile struct {
		Gender int
	}

	var s Student
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if age, ok := field.Tag.Lookup(`li`); ok {
			fmt.Println("explicitly: ", age)
		} else {
			fmt.Println("age = ", age)
		}
	}
}

type Employee struct {
	Id   int64  `liu:"id"`
	Name string `liu:"name"`
	Age  int64  `liu:"age"`
}

func (e *Employee) SetEmployee(name string, id, age int64) error {
	e.Name = name
	e.Id = id
	e.Age = age
	return nil
}
func (e *Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("id = %d, name = %s, age = %d\n", e.Id, e.Name, e.Age)
}

func TestStructMethod(t *testing.T) {
	e := new(Employee)

	et := reflect.TypeOf(e)
	ev := reflect.ValueOf(e)
	numMethod := et.NumMethod()

	t.Logf("numMethod = %d\n", numMethod)

	t.Logf("name = %s, pkgpath = %s, type = %v, func = %v, index = %d\n",
		et.Method(0).Name, et.Method(0).PkgPath, et.Method(0).Type, et.Method(0).Func, et.Method(0).Index)

	t.Log(ev.MethodByName(`SetEmployee`).Call([]reflect.Value{reflect.ValueOf(`allen`), reflect.ValueOf(int64(12)), reflect.ValueOf(int64(20))}))

	info := ev.MethodByName(`GetEmployeeInfo`).Call([]reflect.Value{})
	t.Log(info)

}

func TestStructNameAndSize(t *testing.T) {
	//指针类型
	ep := new(Employee)

	ept := reflect.TypeOf(ep)
	name := ept.Name()
	size := ept.Size()
	kind := ept.Kind()

	t.Logf("pointer: name = %s, size = %d, kind = %v, elem = %s\n", name, size, kind, ept.Elem().String())
	t.Log(ept.String())

	//结构体类型
	var e Employee
	et := reflect.TypeOf(e)
	name = et.Name()
	size = et.Size()
	kind = et.Kind()

	t.Logf("struct: name = %s, size = %d, kind = %s\n", name, size, kind.String())
	t.Log(et.String())

	//string类型
	var s string
	st := reflect.TypeOf(s)
	name = st.Name()
	size = st.Size()
	t.Logf("string: name = %s, size = %d, kind = %s\n", name, size, st.Kind().String())

	//map类型,key为Int64
	var m map[int64]int64
	mt := reflect.TypeOf(m)
	name = mt.Name()
	size = mt.Size()
	t.Logf("map int64: name = %s, size = %d, kind = %s, elem = %s\n", name, size, mt.Kind().String(), mt.Elem().String())

	//map类型，key为string
	var ms map[string]int
	mst := reflect.TypeOf(ms)
	name = mst.Name()
	size = mst.Size()
	t.Logf("map string: name = %s, size = %d, kind = %s, elem = %s\n", name, size, mst.Kind().String(), mst.Elem().String())
	t.Log(mst.String())
}

func TestFieldByNameFunc(t *testing.T) {
	type MaleEmployee struct {
		Employee
		Male bool `liu:"male"`
		age  int  `liu:"a"`
	}

	me := new(MaleEmployee)
	met := reflect.TypeOf(me)
	e := met.Elem()
	sf, ok := e.FieldByNameFunc(func(age string) bool {
		return age == `Age`
	})
	assert.True(t, ok)
	t.Logf("name = %s, tag = %s\n", sf.Name, sf.Tag.Get("liu"))
}
