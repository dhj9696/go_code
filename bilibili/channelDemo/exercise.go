package main

import "fmt"

// 创建一个catChan,最多可以存放10个Cat结构体变量，演示写入和读取的用法
type Cat struct {
	name string
	age  int
}

func main() {
	//定义一个可以存放任意数据类型的管道 3个数据
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"狸花猫", 5}
	allChan <- cat
	//丢弃前面两个数据，得到第三个
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	//下面这两行叫做类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.name=%v\n", a.name)
	//fmt.Printf("newCat.name=%v\n", newCat.name)//错误，编译不通过
	//var catChan chan Cat
	//catChan = make(chan Cat, 10)
	//cat1 := Cat{"Tom1", 10}
	//cat2 := Cat{"Tom2", 20}
	//cat3 := Cat{"Tom3", 30}
	//catChan <- cat1
	//catChan <- cat2
	//catChan <- cat3
	//fmt.Println(cat1, cat2, cat3)

}
