package main

import "fmt"

//一个函数在函数体内又调用了函数本身，我们称为递归调用

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n)
}
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println(n)
	}
}

func main() {
	//test(4)
	test2(4)
}
