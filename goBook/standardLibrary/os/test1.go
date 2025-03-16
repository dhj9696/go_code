package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//使用os.Open()函数打开一个文件
	file, err := os.Open("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// 使用 defer 确保文件在函数结束时关闭
	defer file.Close()
	//创建一个长度为100的byte类型切片
	data := make([]byte, 1024)
	//从文件中读取数据
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	//打印读取的字节数和内容
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
