// map的增删改查
package main

import "fmt"

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("key=", k, "value=", v)
	}
}
func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["中国"] = "北京"
	cityMap["日本"] = "东京"
	cityMap["美国"] = "纽约"
	//遍历
	printMap(cityMap)
	//删除
	delete(cityMap, "日本")
	//修改
	cityMap["美国"] = "河南"
	fmt.Println("-----------------")
	//遍历
	printMap(cityMap)

}
