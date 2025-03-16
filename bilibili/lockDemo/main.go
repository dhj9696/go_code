// 案例要求：计算1-200各个数的阶乘，并且把各个数的阶乘放入到map中
// 最后显示出来，要求使用goroutine来完成
// 思路：
// 1.编写一个函数，来计算各个数的阶乘，并放入到map中
// 2.我们启动的协程多个，统计的结果放入到map中
// 3.map应该做出一个全局的
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁:lock
	//sync是包：synchronized
	lock sync.Mutex
)

// test函数用来计算n的阶乘，并将这个结果放入到myMap中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里我们将res放入到myMap中
	//加锁，防止同时向map中写入数据
	lock.Lock()
	//解锁
	defer lock.Unlock()
	myMap[n] = res
}
func main() {
	//这里开启多个协程
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(3 * time.Second)
	//遍历输出结果
	lock.Lock()
	defer lock.Unlock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
