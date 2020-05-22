package jsonx

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age int64
	Emails []string
}

//总结：Marshal时传递指针，Unmarshal时传递指针的指针；
func TestStruct(t *testing.T) {
	//ex1: p为nil（未初始化），Marshal不会报错，且pBytes为"null"形式的字符串
	var p *Person
	pBytes, err := json.Marshal(p)
	assert.Nil(t, err)
	assert.NotNil(t, pBytes)
	assert.Equal(t, "null", string(pBytes))

	pBytes, err = json.Marshal(&p)
	assert.Nil(t, err)
	assert.NotNil(t, pBytes)
	assert.Equal(t, "null", string(pBytes))


	//ex2: 初始化的struct,各类型取零值，其中空指针时均为字符串null
	p = new(Person)
	pBytes, err = json.Marshal(p)
	assert.Nil(t, err)
	assert.NotNil(t, pBytes)
	assert.JSONEq(t, `{"Name":"","Age":0,"Emails":null}`, string(pBytes))
	assert.Equal(t, `{"Name":"","Age":0,"Emails":null}`, string(pBytes))


	p = new(Person)
	p.Name = "Jackie Chan"
	p.Age = 30
	p.Emails = []string{
		"Jackie@gmail.com",
		"Jacky@gmail.com",
	}

	pBytes, err = json.Marshal(p)
	assert.Nil(t, err)
	t.Log(string(pBytes))

	pBytes, err = json.Marshal(&p)
	t.Log(string(pBytes))

	//ex3: p2未初始化，直接传递p2报错
	var p2 *Person
	err = json.Unmarshal(pBytes, p2)
	assert.NotNil(t, err)
	//t.Error(err)

	err = json.Unmarshal(pBytes, &p2)
	assert.Nil(t, err)
	t.Log(p2)

	// p3初始化了，所以直接传递p3不报错
	p3 := new(Person)
	err = json.Unmarshal(pBytes, p3)
	assert.Nil(t, err)
	t.Logf("p3 = %v\n", p3)
}



type Student struct {
	Name string
	age int64
}

//json只序列化导出的字段
func TestCapitalField(t *testing.T) {
	stu := Student{
		Name: "Jackie",
		age:  22,
	}

	stuBytes, err := json.Marshal(stu)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"Name":"Jackie"}`, string(stuBytes))

	newStu := new(Student)
	err = json.Unmarshal(stuBytes, &newStu)
	assert.Nil(t, err)
	assert.Equal(t, "Jackie", newStu.Name)
	assert.Equal(t, int64(0), newStu.age)
}

// 序列化成小写字段
type  Employee struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}
func TestConvertToLowerField(t *testing.T) {
	employee := new(Employee)
	employee.Id = 2
	employee.Name = "Jackie"

	eBytes, err := json.Marshal(employee)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"id":2,"name":"Jackie"}`, string(eBytes))

	newEmployee := new(Employee)
	err = json.Unmarshal(eBytes, &newEmployee)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), newEmployee.Id)
	assert.Equal(t, "Jackie", newEmployee.Name)
}


// omitempty删除空值字段
type Class  struct {
	Id int64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
func TestOmitEmpty(t *testing.T) {
	class := new(Class)
	classBytes, err := json.Marshal(class)
	assert.Nil(t, err)
	assert.Equal(t, "{}", string(classBytes))

	var newClass *Class
	classBytes, err = json.Marshal(newClass)
	assert.Nil(t, err)
	assert.Equal(t,  "null", string(classBytes))
}

// 跳过字段（及被json忽略的导出字段）
type Teacher struct  {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"-"`
}

type Tutor struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

func TestSkipField(t *testing.T) {
	teacher := new(Teacher)
	teacher.Id =  1001
	teacher.Name  = "Jackie"
	teacher.Age = 23

	tBytes,  err := json.Marshal(teacher)
	assert.Nil(t,  err)
	t.Log(string(tBytes))

	newTeacher := new(Teacher)
	err = json.Unmarshal(tBytes, &newTeacher)
	assert.Nil(t, err)
	assert.Equal(t, int64(1001), newTeacher.Id)
	assert.Equal(t, "Jackie", newTeacher.Name)
	assert.Equal(t, int64(0), newTeacher.Age)

	tutor := new(Tutor)
	err = json.Unmarshal(tBytes, &tutor)
	assert.Nil(t,  err)
}

//Map转换成json跟struct转换后的数据一致，但key的顺序不一定
func TestMapsToJson(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "Jackie"
	m["age"] = 20
	m["emails"] = []string{"allen@gmail.com", "Jackie@gamil.com"}

	mBytes, err := json.Marshal(m)
	assert.Nil(t, err)

	t.Log(string(mBytes))
}


//Slice转化成Json，
func  TestSliceToJson(t  *testing.T) {
	sliceInt := []int64{22, 33}
	sBytes, err := json.Marshal(sliceInt)
	assert.Nil(t, err)
	assert.JSONEq(t, `[22,33]`, string(sBytes))
	t.Log(string(sBytes))

	sliceStr := []string{"22", "33"}
	sbs, err := json.Marshal(sliceStr)
	assert.Nil(t, err)
	assert.JSONEq(t, `["22","33"]`, string(sbs))
	t.Log(string(sbs))
}


// 当不确定Json字段时，应使用Map来接收
func TestJsonToMap(t *testing.T) {
	bytes := []byte(`
		{
			"name": "Jackie",
			"age": 20,
			"emails": ["allen@gmail.com", "jackie@gmail.com"],
			"male": true 
		}
	`)

	js := make(map[string]interface{})
	err := json.Unmarshal(bytes, &js)
	assert.Nil(t, err)

	for key, val := range js {
		t.Logf("%s : %v\n", key, val)
	}

	//注***： 此处go并不知道json里数据类型，所以均以interface{}类型替代，如果[]string会报错；
	emails := js["emails"].([]interface{})

	t.Logf("first email: %s\n", emails[0].(string))
	t.Logf("second email: %s\n", emails[1].(string))
}


// 从json文件中对json流
func TestFileToJson(t *testing.T) {
	fReader, err  := os.Open("jackie.json")
	assert.Nil(t, err)

	person := make([]map[string]interface{}, 0)
	err = json.NewDecoder(fReader).Decode(&person)
	assert.Nil(t, err)

	t.Logf("person = %v\n", person)
}


// 写Json数据至json文件
func TestJsonToFile(t *testing.T) {
	p := []*struct{
		Name string `json:"name"`
		Age int64 `json:"age"`
	}{
		{
			Name: "allen",
			Age: 20,
		},
		{
			Name: "jakie",
			Age: 22,
		},
	}

	fWriter, err := os.Create("output.json")
	assert.Nil(t, err)
	err = json.NewEncoder(fWriter).Encode(p)
	assert.Nil(t, err)
}
