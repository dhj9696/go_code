package main

import (
	"fmt"
	"sync"
)

/*
锁 - 资源竞争
*/
var total int
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁

// 锁不能复制，复制后就失去了锁的效果
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)

}
