package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var myList list.List
	myList := list.New()
	//放在尾部
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	//放在头部
	myList.PushFront("a")
	myList.PushFront("b")
	myList.PushFront("c")
	//遍历打印值,正序
	e := myList.Front()
	for ; e != nil; e = e.Next() {
		if e.Value.(string) == "b" {
			break
		}
	}
	myList.InsertBefore("gin", e)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//反向遍历
	//for e := myList.Back(); e != nil; e = e.Prev() {
	//	fmt.Println(e.Value)
	//}
}
