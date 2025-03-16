package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

func main() {
	/*
		不同类型的数据零值不一样
		bool	false
		numbers 0
		string ""
		pointer nil
		slice nil
		map nil
		channel、interface、function		nil
		struct 默认值不是nil，是具体字段的默认值
	*/
	p1 := Person2{
		name: "bobby",
		age:  10,
	}
	p2 := Person2{
		name: "bobby",
		age:  10,
	}
	if p1 == p2 {
		fmt.Println("p1 and p2 are equal")
	}
	//slice的默认值
	//var ps []Person2
	var ps2 = make([]Person2, 0) //empty slice
	if ps2 == nil {
		fmt.Println("ps2 is nil")
	} else {
		fmt.Println("ps2 is not nil")
	}
	var m = make(map[string]string)
	var m2 = make(map[string]string)
	m["一号选手："] = "jdh"
	m["二号选手："] = "bmh"
	m2["三号选手："] = "张三"
	m2["四号选手："] = "李四"
	for k, v := range m {
		fmt.Println("k:", k, "v:", v)
	}
	for k, v := range m2 {
		fmt.Println("k:", k, "v:", v)
	}
}
