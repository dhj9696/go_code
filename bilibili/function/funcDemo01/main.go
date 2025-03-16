package main

import "fmt"

func test(n1 int) int {
	n1 = n1 + 1
	fmt.Println("test() n1 = ", n1)
	return n1
	//函数在栈帧上存在，与main函数是分离的，
	//所有这里n1的不管怎么变化也与main中的n1没关系
}
func main() {
	n1 := 10
	res := test(n1)
	fmt.Println("main() n1 = ", n1)
	fmt.Println("main() res = ", res)
}
