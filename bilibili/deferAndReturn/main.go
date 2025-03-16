package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer开始执行")
	return 0
}
func returnFunc() int {
	fmt.Println("return开始执行")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
func main() {
	returnAndDefer()
}
