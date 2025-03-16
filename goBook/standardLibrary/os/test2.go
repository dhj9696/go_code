package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	info, err := os.Stat("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件大小: ", info.Size())
	fmt.Println("修改时间: ", info.ModTime())

	//判断是否是，目录
	isDir := info.IsDir()
	fmt.Println(isDir)
}
