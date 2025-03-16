package main

import (
	"fmt"
	"time"
)

// channel是多线程安全，要初始化

func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}
func g2(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
}
func main() {
	//select 类似于switch case语句，但是select的功能和Linux里面提供的io的select、poll、epoll
	//select 主要作用于多个channel
	//现在有个需求，有两个goroutine都在执行，但是我在主的goroutine中，当某一个执行完成后，这个时候我会立马知道
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go g1(g1Channel)
	go g2(g2Channel)

	//我要监控多个channel，任何一个channel返回值都知道
	//1.某一个分支就绪了就执行该分支 2.如果两个都就绪了，随机执行
	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("timeout")
			return
		}
	}
}
