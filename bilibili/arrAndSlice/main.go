package main

import "fmt"

/*
数组：数组的长度是固定的，在定义时需要指明数组的长度
切片：切片是动态的长度，不需要指定长度。可以通过make函数创建，也可以基于数组或切片创建
长度特性
数组的长度是固定的，一旦定义，其长度不能改变
切片的长度是动态的，可以通过append函数向切片中添加元素，从而改变其长度
*/
func main() {
	//定义一个长度为5的整数数组
	var arr [5]int
	fmt.Println(arr)
	//使用make函数创建一个初始长度为3，容量为5的整数切片
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	//基于数组创建切片
	arr1 := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := arr1[1:3]
	fmt.Println(sliceFromArr)
}
