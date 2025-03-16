package main

import "fmt"

// 定义一个无参数、无返回值的函数
func sayHello() {
	fmt.Println("hello world")
}

// 定义一个有参数、无返回值的函数
func printSum(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

// 定义一个有参数有返回值的函数(需要return)
func add(a, b int) int {
	return a + b
}

// 定义一个有多个返回值的函数
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 定义一个使用命名返回值的函数
func calculate(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return
}
func main() {
	//调用函数
	sayHello()
	printSum(3, 5)
	//调用函数并接收返回值
	result := add(5, 6)
	fmt.Println(result)
	//调用函数并接收多个返回值
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	s, d := calculate(11, 6)
	fmt.Println(s, d)

}
