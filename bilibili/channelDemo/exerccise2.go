package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//创建一个channel
	ch := make(chan Person)
	//设置随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//启动一个goroutine来生成并发送Person实例到channel
	go func() {
		defer close(ch)
		names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}
		addresses := []string{"CityA", "CityB", "CityC", "CityD", "CityE"}
		for i := 0; i < 10; i++ {
			p := Person{
				Name:    names[r.Intn(len(names))],
				Age:     r.Intn(100),
				Address: addresses[r.Intn(len(addresses))],
			}
			ch <- p
		}
	}()
	//遍历channel并打印Person实例信息
	for p := range ch {
		fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
	}
}
