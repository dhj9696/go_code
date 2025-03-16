package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for i = 0; i < 20; i++ {
		fmt.Printf("new Goroutin: i =%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	//创建一个go线程，去执行newTask() 流程
	go newTask()
	//等待20秒，确保子goroutine执行完
	time.Sleep(20 * time.Second)
}
