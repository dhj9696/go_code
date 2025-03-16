package main

import "fmt"

func main() {
	//go折中，slice 切片
	//var course []string
	//fmt.Printf("%T\n", course)
	//向切片添加数据
	//course = append(course, "go")
	//course = append(course, "grpc")
	//fmt.Println(course)

	//切片的初始化
	//1.基本初始化：初始化一个包含五个元素的切片
	//numbers1 := []int{1, 2, 3, 4, 5}
	//fmt.Println(numbers1)
	////2.使用make函数初始化
	////make([]T, length, capacity)
	////创建一个长度为3，容量为5的int类型切片
	//numbers2 := make([]int, 3, 5)
	//fmt.Printf("切片类型:%T\n", numbers2)
	//fmt.Println("长度:", len(numbers2))
	//fmt.Println("容量", cap(numbers2))
	////3.从数组或切片中截取初始化
	////初始化一个数组
	//array := [5]int{1, 2, 3, 4, 5}
	////从数组中截取一个切片
	////array[1:3]表示从数组array中截取索引从1到2的元素，创建一个新的切片slice
	//slice := array[1:3] //1为起始索引（包含），3是结束索引（不包含）
	//fmt.Println(slice)
	////4.初始化空切片
	////方式一：使用字面量
	//var slice1 []int
	//fmt.Println("slice1", slice1)
	////方式二：使用make函数
	//slice2 := make([]int, 0)
	//fmt.Println("slice2:", slice2)
	//
	////访问切片的元素
	////访问单个
	//courseSlice := []string{"go", "grpc", "gin", "mysql", "redis"}
	//fmt.Println(courseSlice[1])
	////访问多个
	//fmt.Println(courseSlice[1:4]) //访问第2（包含）到第5个元素（不包含），左开右闭
	//fmt.Println(courseSlice[1:])  //访问第2个到结尾的元素（包含）
	//fmt.Println(courseSlice[:3])  //访问第4个到开始的元素（不包含）
	//fmt.Println(courseSlice[:])   //访问所有元素

	//courseSlice := []string{"go", "grpc"}
	//courseSlice2 := []string{"mysql", "es", "gin"}
	////courseSlice = append(courseSlice, "gin", "mysql", "es")
	////将courseSlice2中从第2个元素开始的所有元素追加到courseSlice切片中
	//courseSlice = append(courseSlice, courseSlice2[1:]...) //有"..."才可以
	//fmt.Println(courseSlice)

	//如何删除slice中的元素
	courseSlice := []string{"go", "grpc", "gin", "mysql", "redis"}
	//不能直接删除，只有分步截取
	mySlice := append(courseSlice[:2], courseSlice[3:]...)
	//fmt.Println(mySlice)

	//复制slice
	var mySliceCopy = make([]string, len(courseSlice))
	copy(mySliceCopy, mySlice)
	fmt.Println(mySliceCopy)

	mySliceCopy2 := mySlice[:]
	fmt.Println(mySliceCopy2)

	fmt.Println("----------------")
	mySlice[0] = "java"
	fmt.Println(mySliceCopy)
	fmt.Println(mySliceCopy2)
}
