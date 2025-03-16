package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串的比较
	a := "Hello"
	b := "world"
	fmt.Println(a != b)

	fmt.Println(a > b)

	name := "我是jdh*hhh"
	//name中是否包含“我是”
	re := strings.Contains(name, "我是")
	fmt.Println(re)
	//name中h出现的次数
	fmt.Println(strings.Count(name, "h"))
	//分割字符串
	fmt.Println(strings.Split(name, "*"))
	//判断字符串的开头
	fmt.Println(strings.HasPrefix(name, "我"))
	//判断字符串的结尾
	fmt.Println(strings.HasSuffix(name, "*"))
	//查找指定字符的出现位置
	fmt.Println(strings.Index(name, "j"))
	//替换指定字符
	fmt.Println(strings.Replace(name, "jdh", "bmh", -1))

}
