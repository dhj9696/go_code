package utils

import "fmt"

// Cal 将计算的功能，放到一个函数中，然后在需要使用是调用即可
// 不成文的规定：为了让其它包的文件使用Cal函数，需要将函数名首字母需要大写，小写表示私有
func Cal(n1 float64, n2 float64, operator string) float64 {
	var res float64
	switch operator {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	default:
		fmt.Println("操作符错误")
	}
	return res

}
