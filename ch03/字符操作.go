package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//长度计算
	str := "我是jdh"
	byteLength := len(str)
	fmt.Println(byteLength)

	name := "我是jdh"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义符
	courseName := "hello\r\ngo体系课程"
	fmt.Println(courseName)
	fmt.Print("hello\r\n")
	fmt.Print("world\r\n")

	//格式化输出
	username := "jdh"
	age := 21
	address := "北京"
	mobile := "121324125"
	fmt.Println("用户名：" + username + ", 年龄：" + strconv.Itoa(age) + ", 地址：" + address + ", 电话：" + mobile) //不常用，不易维护，但是性能高
	fmt.Printf("用户名：%s, 年龄：%d, 地址：%s, 电话：%s\r\n", username, age, address, mobile)                       //常用，易维护，但是性能低
	userMsg := fmt.Sprintf("用户名：%s, 年龄：%d, 地址：%s, 电话：%s\r\n", username, age, address, mobile)
	fmt.Print(userMsg)

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString(", 年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(", 地址：")
	builder.WriteString(address)
	builder.WriteString(", 电话：")
	builder.WriteString(mobile)
	re := builder.String()
	fmt.Println(re)
}
