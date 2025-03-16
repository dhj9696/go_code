package main

import (
	"fmt"
	"time"
)

// python,java,php 多线程编程、多进程编程,多线程和多进程存在的问题主要是耗费内存
// 内存、线程切换，web2.0，
// go协程-goroutine

func main() {
	//匿名函数启动goroutine

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i) //值传递
	}
	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
