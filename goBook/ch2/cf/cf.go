// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"go_code/goBook/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	//从键盘获取输入
	//var t int
	//fmt.Print("请输入你要转换的温度：")
	//fmt.Scanln(&t)

	//以命令行的方式获取输入
	for _, arg := range os.Args[1:] {
		t, _ := strconv.ParseFloat(arg, 64)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		//	os.Exit(1)
		//}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %v, %s = %v\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}
