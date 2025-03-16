package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//process, err := os.FindProcess(os.Getpid())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("进程名: ", process.Name())
	//获取当前进程id
	pid := os.Getpid()
	fmt.Println("当前进程ID: ", pid)
	//获取当前进程的科执行文件路径
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("获取可执行文件路径时出错:", err)
		return
	}
	//从路径中提取进程名称
	processName := filepath.Base(executablePath)
	fmt.Println("进程名: ", processName)
}
