package main

import (
	"fmt"
	"strings"
)

/*
案例
1)编写一个函数makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg),
并返回一个闭包
2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg)，则返回文件名.jpg，
如果已经有.jpg后缀，则返回原文件名.
3)要求使用闭包的方式完成
4)strings.HasSuffix，该函数可以判断某个字符串是否有指定的后缀
*/
func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果name没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	//测试一下makeSuffix的使用
	//返回一个闭包
	f := makeSuffix(".jpg")
	fmt.Println("文件处理后=", f("winter"))
	fmt.Println("文件处理后=", f("bird.jpg"))

}
