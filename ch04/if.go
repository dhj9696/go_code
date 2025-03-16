package main

import "fmt"

/*
	if 布尔表达式 {
		逻辑
	}
*/
func main() {
	age := 17
	country := "美国"
	if (age < 18) && (country == "中国") {
		fmt.Println("未成年")
	} else {
		fmt.Println("成年了")
	}
}
