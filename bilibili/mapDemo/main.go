package main

import "fmt"

func main() {
	//第一种声明方式
	//声明myMap1是一种map类型 key是string, value是string
	var myMap1 map[string]string
	//在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	//第三种声明方式
	myMap3 := map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	fmt.Println(myMap3)
}
