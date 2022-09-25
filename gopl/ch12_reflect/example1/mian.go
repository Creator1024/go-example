package main

import (
	"fmt"
	"reflect"
)

type MyInt int64

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	mi := reflect.TypeOf(MyInt(1))
	fmt.Printf("%v, %v\n", mi.Name(), mi.Kind())
	var i = int64(1)
	pi := reflect.TypeOf(&i)
	fmt.Printf("%v, %v\n", pi.Name(), pi.Kind()) // , ptr

	stu1 := student{
		Name:  "stu1",
		Score: 100,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct

	// 遍历每个结构体字段，获取相关信息
	// name:Name index:[0] type:string json tag:name
	// name:Score index:[1] type:int json tag:score
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	// name:Score index:[1] type:int json tag:score
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
