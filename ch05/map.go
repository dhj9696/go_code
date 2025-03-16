package main

import "fmt"

func main() {
	//map是一个key(索引)和value(值)的无序集合，主要是查询方便
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}
	//fmt.Println(courseMap["grpc"]) //根据索引寻找值
	courseMap["mysql"] = "mysql原理" //存值(需初始化)
	courseMap["java"] = "我是java"
	//fmt.Println(courseMap)
	//map必须初始化才能使用 1.map[string]string{} 2.make(map[string]string, 3)
	//var courseMake = make(map[string]string, 3) //make函数主要用于初始化slice、map、channel
	//fmt.Println(courseMake)
	//遍历
	for k, v := range courseMap {
		fmt.Println(k, v)
	}
	//d, ok := courseMap["java"]
	//if ok {
	//	fmt.Println("find", d)
	//} else {
	//	fmt.Println("not find")
	//}
	if d, ok := courseMap["mysql"]; ok {
		fmt.Println(d)
	} else {
		fmt.Println("not find")
	}
	//删除一个元素
	delete(courseMap, "grpc")
	fmt.Println(courseMap)
}
