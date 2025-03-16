package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int        //定义一个int类型的channel(通道) msg 并未初始化
	msg = make(chan int, 2) //将上一行定义的msg使用make初始化为容量为2的缓冲通道
	//msg := make(chan int, 2 ) 也可以这样定义加初始化一步成
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)
	msg <- 1 //放值到channel中
	msg <- 2
	msg <- 3
	close(msg)
	time.Sleep(time.Millisecond * 3)
}
