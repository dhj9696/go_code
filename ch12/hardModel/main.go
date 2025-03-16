package main

//多个生产者和多个消费者模型
import (
	"fmt"
	"sync"
)

// 生产者函数
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- id*10 + i
		fmt.Printf("Producer %d is produced: %d\n", id, id*10+i)
	}
}

// 消费者函数
func consumer(id int, ch <-chan int, wg *sync.WaitGroup, finish chan struct{}) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Consumer %d is consumed: %d\n", id, num)
	}
	finish <- struct{}{} //通知主函数该消费者已完成
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	finish := make(chan struct{})
	//启动两个生产者
	numProducers := 2
	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(i, ch, &wg)
	}
	//启动3个消费者
	numConsumers := 3
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consumer(i, ch, &wg, finish)
	}
	//等待所有生产者完成
	go func() {
		wg.Wait()
		close(ch)
	}()
	//等待所有消费者完成
	for i := 0; i < numConsumers; i++ {
		<-finish
	}
	close(finish)
	fmt.Println("All tasks are done")
}
