package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(i)
	//}

	//求1+2+...+100
	//var sum int
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//打印99乘法表
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d = %-4d", j, i, j*i)
	//	}
	//	fmt.Println()
	//}

	//name := "imooc go"
	//for _, value := range name {
	//	//fmt.Println(index, value)
	//	fmt.Printf("%c\n", value)
	//}

	name := "imooc go体系课"
	nameRune := []rune(name)
	print(len(name))
	//遍历name
	for i := 0; i < len(nameRune); i++ {
		fmt.Printf("%c\r\n", nameRune[i])
	}
}
