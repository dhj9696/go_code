package main

import (
	"fmt"
	"strconv"
)

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str = str + strconv.Itoa(x)
		fmt.Println("str = ", str)
		return n
	}
	//返回的是一个匿名函数，但是这个匿名函数引用到函数外的n,
	//因此这个匿名函数就和n形成一个整体，构成的整体就是闭包
}
func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
