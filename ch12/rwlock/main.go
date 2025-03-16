package main

import (
	"fmt"
	"sync"
	"time"
)

//锁本质上是将并行的代码串行化了，会影响性能（下降）
//即便是设计锁，那么也应该尽量的保证并行
//有两组协程，一组负责写数据，另一组负责读数据，web系统读多写少
//虽然有多个goroutine,但是读协程之间应该并发，读和写之间烟应该串行
//读写锁

func main() {

	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(6)
	//启动一个匿名的goroutine并执行
	//写的goroutine
	go func() {
		time.Sleep(3 * time.Second)
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁获取和读锁获取
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(5 * time.Second)

	}()

	//读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() //加读锁，读锁不会阻止别人的读
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
		}()
	}
	wg.Wait()
}
