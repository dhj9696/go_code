package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统cpu的数量
	num := runtime.NumCPU()
	fmt.Printf("NumCPU: %d\n", num)
	//可以自主设置使用多少个cpu
	runtime.GOMAXPROCS(num - 1)
}
