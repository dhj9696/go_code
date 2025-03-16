package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	//创建一个命令对象
	cmd := exec.Command("type")
	//创建输入管道
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("创建输入管道失败: ", err)
		return
	}
	//创建输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("创建输出管道失败: ", err)
		return
	}
	//启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("启动command失败: ", err)
		return
	}
	//向子进程发送数据
	message := "你好， 孩子进程!"
	_, err = stdin.Write([]byte(message))
	if err != nil {
		fmt.Println("向输入端写入数据失败: ", err)
		return
	}
	stdin.Close()

	//从子进程读取数据
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println("接收到了孩子进程", scanner.Text())
	}
	//等待命令执行完毕
	err = cmd.Wait()
	if err != nil {
		fmt.Println("等待command错误", err)
		return
	}
}
