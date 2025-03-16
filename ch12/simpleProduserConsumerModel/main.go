package main

import (
	"fmt"
	"sync"
)

// 简单的生产者-消费者模型
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Producer: %d\n", i)
	}
	close(ch) //生产结束后关闭通道
}

// 消费者函数
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Consumer: %d\n", num)
	}
}
func main() {
	var wg sync.WaitGroup
	//没有指定容量-属于无缓冲通道
	ch := make(chan int)
	//启动生产者
	wg.Add(1)
	go producer(ch, &wg)
	//启动消费者
	wg.Add(1)
	go consumer(ch, &wg)
	//等待所有goroutine完成
	wg.Wait()
	fmt.Println("All tasks are done.")
}
