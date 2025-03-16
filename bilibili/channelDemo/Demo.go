package main

import "fmt"

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//2.看一下intChan是什么
	fmt.Printf("intChan的值=%v\nintChan本身的地址=%p\n", intChan, &intChan)
	//3.向管道写入数据
	intChan <- 10
	num := 985
	num3 := 211
	//给管道写入的数据不能超过其容量
	intChan <- num
	intChan <- num3
	<-intChan //取出管道中的数据，并丢掉
	//4.看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
	//管道最大的作用就是一边写入数据，一边读出数据，就像现实生活中的管道一样，一个口流入（写入），一个口流出（读出）

}
