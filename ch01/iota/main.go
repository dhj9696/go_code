package main

import "fmt"

func a() (int, bool) {
	return 0, true
}

func main() {
	//匿名变量
	var _ int
	_, ok := a()
	if ok {
		fmt.Println("ok")
	}

	//iota，特殊常量，可以认为是一个可以被编译器修改的常量
	const (
		ERR1 = iota + 1
		ERR2
		ERR25 = "haha"
		ERR3
		ERR4 = iota
	)
	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR4)
}
