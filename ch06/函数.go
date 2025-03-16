package main

import (
	"fmt"
	"time"
)

func add(items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return sum, err
}
func add2(a, b int) {
	fmt.Printf("sum is: %d\n", a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}

func cal(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("啥也不是")
		}
	}
}

func runForever() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("doing")
	}
}

var local int

func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	} //闭包
}

func main() {
	/*
		go中函数是"一等公民"
		1.函数本身可以当作变量
		2.匿名函数 闭包
		3.函数可以满足接口
	*/
	funcVar := add

	a := 1
	b := 2
	sum, _ := funcVar(a, b, 3, 4)
	fmt.Println(sum)
	fmt.Println(a)

	cal("+")()

	callback(1, add2)

	//匿名函数
	localFunc := func(a, b int) {
		fmt.Printf("total is: %d\n", a+b)

	}
	callback(1, localFunc)

	//闭包
	//函数每调用一次返回值加一
	nextFunc := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
	nextFunc2 := autoIncrement()
	for i := 0; i < 10; i++ {
		fmt.Println(nextFunc2())
	}
}
