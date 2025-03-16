package main

import "fmt"

func main() {
	//常量， 定义的时候指定值，不能修改
	const PI float32 = 3.1415926 //显示定义
	const PI2 = 3.1415926        //隐式定义
	const (
		x int = 16
		y
		s = "hello"
		z
	)
	fmt.Println(x, y, s, z)
}
