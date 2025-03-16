package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}
func mPrint2(data interface{}) {
	fmt.Println(data)
}
func main() {
	//var data = []interface{}{
	//	"bobby", 18, 1.80,
	//}
	//mPrint(data...)
	var data = []string{
		"bobby", "bobby2", "bobby3",
	}
	var datai []interface{}
	for _, value := range data {
		datai = append(datai, value)
	}
	mPrint(datai...)
}
