package main

import "fmt"

var (
	//fun1就是一个全局匿名函数
	Fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//匿名函数使用方式1
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	//案例演示，求两个数的和，使用匿名函数的方式完成
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("result1 = ", res1)
	//匿名函数使用方式2
	//将匿名函数func(n1, n2 int)int付给a变量
	//则a的数据类型就是函数类型，此时我们可以通过a完成调用
	a := func(n1, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println("result2 = ", res2)
	//全局匿名函数的使用
	res3 := Fun1(10, 20)
	fmt.Println("result3 = ", res3)
}
