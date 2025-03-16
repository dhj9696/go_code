package main

import (
	"fmt"
	"time"
)

/*
默认情况下，channel是双向的
但是，我们经常将一个channel作为参数进行传递，希望对方是单向使用，只能读不能写
*/
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
}

func main() {
	//var ch1 chan int       //双向channel
	//var ch2 chan<- float64 //单向channel，只能写入float64的数据
	//var ch3 <-chan int     //单向channel，只能读取数据
	//c := make(chan int, 3)  //初始化一个双向的channel
	//var send chan<- int = c //将c改成单向的channel，send-only
	//var read <-chan int = c //将c改成单向的channel，read-only
	//send <- 1
	//<-read

	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(10 * time.Second)
}
