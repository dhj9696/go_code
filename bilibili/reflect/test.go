package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is calling")
	fmt.Printf("%v\n", u)
}
func main() {
	user := User{1, "jdh", 21}
	DoFileAndMethod(user)
}
func DoFileAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType=", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue=", inputValue)

	//通过type获取里面的字段
	//1. 获取interface的reflect.Type，通过Type的NumField进行遍历

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	//通过Type获取里面的方法调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}
