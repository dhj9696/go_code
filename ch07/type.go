package main

import "fmt"

// 1.定义自定义类型
// 定义一个自定义类型Celsius表示摄氏度
//type Celsius float64

// Fahrenheit 定义一个自定义类型Fahrenheit表示华氏度
// type Fahrenheit float64
//
// // CTof 定义一个将摄氏度转换为华氏度的函数
//
//	func CTof(c Celsius) Fahrenheit {
//		return Fahrenheit(c*9/5 + 32)
//	}
func main() {
	//	//type关键字
	//	/*
	//		1.定义结构体
	//		2.定义接口
	//		3.定义类型别名
	//	*/
	//	var c Celsius = 27
	//	f := CTof(c)
	//	fmt.Printf("%g摄氏度是%g华氏度\n", c, f)
	//
	//}

	// 2.定义结构体
	type Person struct {
		name    string
		age     int
		address string
		height  float32
	}
	//
	//func main() {
	//	//如何初始化结构体
	//	p1 := Person{
	//		name:    "jdh",
	//		age:     21,
	//		address: "漯河",
	//		height:  1.70,
	//	}
	//	var person []Person
	//	person = append(person, p1)
	//	fmt.Println(person)
	//}

	// 3.自定义接口
	// 定义一个接口类型Shape表示形状
	//type Shape interface {
	//	Area() float64
	//}
	//
	//// 定义一个结构体类型Rectangle表示矩形
	//type Rectangle struct {
	//	Width, Height float64
	//}
	//
	//// 为Rectangle 类型实现Shape接口的Area方法
	//func (r Rectangle) Area() float64 {
	//	return r.Width * r.Height
	//}
	//func main() {
	//	r := Rectangle{Width: 10, Height: 20}
	//	var s Shape = r
	//	fmt.Printf("矩形的面积是: %g\n", s.Area())

	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京",
		"海淀区",
		"海淀小区",
	}
	fmt.Println(address.city)
}
