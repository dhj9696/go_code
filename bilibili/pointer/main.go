package main

import "fmt"

func swapValue(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
func swapValuePointer(pa, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb //通过调用main函数传递参数的地址对其进行修改
	*pb = temp
}
func main() {
	a := 10
	b := 20
	c := 30
	d := 40
	swapValue(c, d) //以值的形式传递给swap，a,b在函数调用结束之后就销毁了，main函数中的a,b的值没有改变
	swapValuePointer(&a, &b)
	fmt.Println(c, d)
	fmt.Println("a = ", a, "b = ", b)
}

//func changeValue(p *int) {
//	*p = 10
//	fmt.Printf("%p\n", p)
//}
//func main() {
//	a := 1
//	changeValue(&a)
//	fmt.Println("a = ", a)
//}
