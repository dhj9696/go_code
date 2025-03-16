package main

import "fmt"

// 数组定义：var name [count]int
func main() {
	//var courses1 [3]string //courses，只有3个元素的数组
	//var courses2 [4]string //courses，有4个元素的数组
	//fmt.Printf("%T\r\n", courses1)
	//fmt.Printf("%T\r\n", courses2)

	//courses1[0] = "Go"
	//courses1[1] = "C++"
	//courses1[2] = "C#"
	//fmt.Println(courses1)
	//
	//for _, value := range courses1 {
	//	fmt.Println(value)
	//}
	//数组的初始化
	//var courses1 = [3]string{"go", "grpc", "gin"}
	//for _, value := range courses1 {
	//	fmt.Println(value)
	//}
	//var courses2 = [3]string{2: "gin"}
	//for _, value := range courses2 {
	//	fmt.Println(value)
	//}
	//courses3 := [...]string{"go", "grpc"}
	//遍历方法1
	//for _, value := range courses3 {
	//	fmt.Println(value)
	//}
	//遍历方法2
	//for i := 0; i < len(courses3); i++ {
	//	fmt.Printf("%s ", courses3[i])
	//
	//}
	//fmt.Println()
	//courses4 := [...]string{"go", "grpc"}
	//if courses3 == courses4 {
	//	fmt.Println("equal")
	//}

	//多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "bobby", "体系课"}
	//fmt.Println(courseInfo[0])
	courseInfo[1] = [4]string{"grpc", "2h", "bobby2", "grpc入门"}
	courseInfo[2] = [4]string{"gin", "1.5h", "bobby3", "gin高级开发"}
	//for i := 0; i < len(courseInfo); i++ {
	//	for j := 0; j < len(courseInfo[i]); j++ {
	//		fmt.Printf("%s ", courseInfo[i][j])
	//	}
	//	fmt.Println()
	//}
	for _, row := range courseInfo {
		fmt.Println(row)
	}

}
