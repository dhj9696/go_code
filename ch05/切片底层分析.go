package main

import "fmt"

func main() {
	var data []int
	for i := 0; i < 2000; i++ {
		data = append(data, i)
		fmt.Printf("len: %d, cap: %d\n", len(data), cap(data))
	}
	//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := data[1:6]
	//s2 := data[2:7]
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(len(s2), cap(s2))
	//fmt.Println(s1, s2)
}
