package main

import (
	"fmt"
	"runtime"
	"time"
)

// 用go创建承载一个形参为空，返回值为空的一个函数\
func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine
			runtime.Goexit()
			defer fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
