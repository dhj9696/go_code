package main

import "fmt"

type Person struct {
	name string
}

// 接收者
func (pn *Person) SayHello() {

}

// 通过指针交换两个值
func swap(a, b *int) {
	//临时值
	t := *a
	//将a地址的值改为b指向的值
	*a = *b
	//将b地址的值改为a指向的值
	*b = t

}
func changeName(p *Person) {
	p.name = "imooc"
}
func main() {
	//指针
	//这个提个需求，我希望结构体传值的时候，我在函数中修改的值可以反应到变量中
	//p1 := Person{
	//	name: "bobby",
	//}
	//var ptr *Person = &p1
	//fmt.Printf("ptr: %p\n", ptr)
	////changeName(&p)
	//fmt.Println(p.name)

	//指针的定义
	//po := &Person{
	//	name: "bobby2",
	//}
	//po.name = "bobby4"
	//fmt.Println(po)

	//var a int = 10
	//b := &a
	//fmt.Println(b)
	//指针需要初始化
	//var b *int
	//var p *Person
	//fmt.Println(p.name)//p为初始化，是一个空指针，运行时会报错
	//fmt.Println(b)
	//ps := &Person{} //第一种初始化方式
	//var emptyPerson Person
	//pi := &emptyPerson //第二种初始化方式
	//var pp = new(Person) //第三种初始化方式 (推荐)
	//fmt.Println(pp.name)
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
