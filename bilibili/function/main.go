package main

import (
	"fmt"
	"go_code/bilibili/utils"
)

func main() {
	num1 := 3.4
	num2 := 3.4
	operator := "+"
	re := utils.Cal(num1, num2, operator)
	fmt.Printf("result = %.2f\n", re) //保留两位小数
}
