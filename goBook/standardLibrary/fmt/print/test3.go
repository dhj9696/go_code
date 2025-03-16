package main

import (
	"fmt"
	"os"
)

/*
定义了常量name为 “Kim”，age为 22。
接着使用fmt.Fprint函数将字符串 “Kim is 22 years old.”
写入标准输出，并将返回的字节数和可能出现的错误分别赋值给变量n和err。
如果发生错误，就使用fmt.Fprintf函数将错误信息写入标准错误输出。
最后，打印出写入的字节数和提示信息 “bytes written.”。
*/
func main() {
	const name, age = "kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")
	//Fprint 的 n 和 err 返回值是
	// 底层 io. Writer 返回的那些。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Println(n, " bytes written ")

}
