package main

import (
	"fmt"
	"time"
)

/*
go语言线程(goroutine)之间的通信原则：不要通过共享内存来通信，而要通过通信来实现内存共享
通信机制：消息队列。如消费者与生产者之间的关系：消费者生成的数据放入消息队列中，消费者去生产队列中取
通信使用函数:channel
*/
func main() {
	//无缓冲channel适用于通知，B要第一时间知道A是否已经完成
	//有缓冲channel适用于消费者和生产者之间的通信
	/*
		go中channel的应用场景
		1.消息传递、消息过滤
		2.信号广播
		3.事件订阅和广播
		4.任务并发
		5.结果汇总
		6.并发控制
		7.同步和异步
	*/
	var msg chan string
	var msg2 chan string
	//有缓冲通道
	msg = make(chan string, 1)
	msg <- "hello jdh"
	data := <-msg
	fmt.Println(data)
	//无缓冲通道
	msg2 = make(chan string, 0)
	go func(msg2 chan string) {
		data := <-msg2
		fmt.Println(data)
	}(msg2)
	msg <- "hello bmh" //放值到channel中
	time.Sleep(time.Second)

}
