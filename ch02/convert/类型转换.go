package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var a int8 = 12
	//var b uint8 = uint8(a)
	//
	//var f float32 = 3.14
	//var c = int32(f)
	//fmt.Println(a, b, c)

	//字符串转数字
	var isstr = "12"
	myint, err := strconv.Atoi(isstr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	var myi = 32
	fmt.Println(strconv.Itoa(myi))

	//字符串转换为float32,转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("-32", 10, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	parseBool, err := strconv.ParseBool("false")
	if err != nil {
		return
	}
	fmt.Println(parseBool)

	//基本类型转字符串
	boolStr := strconv.FormatBool(true)
	fmt.Println(boolStr)
}
