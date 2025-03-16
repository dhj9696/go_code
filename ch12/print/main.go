package main

import (
	"fmt"
	"time"
)

/*
使用两个goroutine交替打印序列，一个goroutine打印数字，另一个打印字母：
12AB34CD56EF78GH910IJ..............
*/
var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		//我该怎么做，应该此处，等待另一个goroutine来通知我
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		//我该怎么做，应该此处，等待另一个goroutine来通知我
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Printf(str[i : i+2])
		i += 2
		number <- true
	}
}

func main() {
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(100 * time.Second)
}
