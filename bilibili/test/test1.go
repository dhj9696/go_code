package main

import "fmt"

func main() {
	var a string
	a = "aceld"
	//pair<type:string, value:"aceld">
	var allType interface{}
	allType = a
	str, ok := allType.(string)
	fmt.Println(str, ok)
}
