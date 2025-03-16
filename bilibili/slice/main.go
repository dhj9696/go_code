package main

import "fmt"

// 切片的概念：可变的数组
func main() {
	numbers := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//向numbers切片中追加一个元素1，numbers len = 4, [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	//var intArr = [...]int{1, -5, 39, 100, 90}
	//fmt.Println(intArr)
	//var slice = make([]int, 4, 5)
	//slice[1] = 123240
	//slice[3] = 20
	//fmt.Println(slice)
	//slice2 := []int{1, 2, 3}
	//capacity := cap(slice2)
	//fmt.Printf("切片的容量是：%d\n", capacity)
}
