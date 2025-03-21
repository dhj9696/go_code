package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func main() {
	// 定义一个测试用的 uint64 类型的整数
	x := uint64(0xFFFFFFFFFFFFFFFF)
	// 调用 PopCount 函数计算 x 的二进制表示中 1 的个数
	result := PopCount(x)
	// 输出结果
	fmt.Printf("The population count of %d is %d\n", x, result)
}
