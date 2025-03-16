package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	return 0, errors.New("this is an error")
}

func main() {
	//连接数据库、打开文件、上锁，最后都需要关闭数据库、关闭文件、解锁使用error用来防止遗忘
	if _, err := A(); err != nil {
		fmt.Println(err)
	}

}
