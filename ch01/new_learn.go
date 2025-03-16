package main

import "fmt"

// 全局变量
var (
	name1 = "body1"
	name2 = "body2"
	name3 = "body3"
)

func main() {
	//1.变量必须先定义后使用 2.变量必须有类型
	//定义变量的常用方式：
	//var name string
	//age := 1(这种定义方式属于简洁定义，不允许在全局变量中定义)
	age := 21 //在go语言中局部变量定义了不使用是不行的
	fmt.Println("你的年龄是：", age)

	//定义多个变量
	var user1, user2, user3 = "bobby1", "bobby2", "bobby3"
	fmt.Println(user1, user2, user3)

}
