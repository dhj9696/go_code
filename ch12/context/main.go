package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 新的需求：可以主动退出监控程序
// 共享变量，不如消息管道

func cpuInfo(ctx context.Context) {
	//这里拿到一个请求的id
	fmt.Printf("traceid: %v\n", ctx.Value("traceid"))
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu的信息")
		}

	}
}
func main() {
	//有一个goroutine监视cpu的信息
	wg.Add(1)
	//context包提供了三种函数，WithCancel,WithTimeout,WithValue
	//1.ctx, cancel := context.WithCancel(context.Background())

	//2.timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	//3.WithDeadline 在时间点cancel

	//4.WithValue
	valueCtx := context.WithValue(ctx, "traceid", "wq12fas")
	go cpuInfo(valueCtx)
	wg.Wait()
	fmt.Println("监控完成")
}
