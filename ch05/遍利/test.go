package main

import "fmt"

func main() {
	//numbers := []int{1, 2, 3, 4, 5}
	//for key, value := range numbers {
	//	fmt.Println(key, value)
	//}
	// 定义一个映射
	person := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Cathy": 35,
	}
	for key, value := range person {
		fmt.Printf("%s is %d years old.\n", key, value)
	}
}
