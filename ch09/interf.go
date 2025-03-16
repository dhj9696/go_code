package main

import "fmt"

// 接口的定义
type Duck interface {
	// 方法的申请
	Gaga()
	Walk()
	Swim()
}
type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")

}
func (pd *pskDuck) Walk() {
	fmt.Println("鸭子漫步")
}
func (pd *pskDuck) Swim() {
	fmt.Println("鸭子游泳")
}
func main() {
	//go语言的接口，鸭子类型
	var d Duck = &pskDuck{}
	d.Gaga()
	d.Walk()
	d.Swim()

}
