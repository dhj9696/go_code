/*
案例：
1.在主线程(可以理解成进程)中，开启一个goroutine，该协程每隔1秒输出"hello,world"
2.在主线程中也每隔一秒输出"hello,jdh"，输出10次后，退出程序
3.要求主线程和goroutine同时执行
*/
//编写一个函数，每隔一秒输出“hello,jdh”

package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 5; i++ {
		fmt.Println("test()" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go test() //开启一个协程

	for i := 1; i <= 5; i++ {
		fmt.Println("main()" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
