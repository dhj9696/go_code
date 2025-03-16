package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc 被调用")
	fmt.Println(arg)

	//interface如何区分 此时引用的底层数据类型到底是什么？
	//答案
	//给interface{} 提供类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string")
	} else {
		fmt.Println("arg is string type, value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	author string
}

func main() {
	//book := Book{"Golang"}
	//myFunc(book)
	myFunc(11)

}
