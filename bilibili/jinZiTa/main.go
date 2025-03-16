package main

import "fmt"

// 使用for循环完成下面的案例请编写一个程序，可以接收一个整数，表示层数，打印金字塔
// 编程思路
// 1.打印一个矩形
/*
*****
*****
*****
*****
*****
 */
//2.打印半个金字塔
/*
*
**
***
****
*****
 */
//3.打印整个金字塔
/*
    *     1层1个*  空格 4
   ***	   2层3个* 空格 3
  *****   3层5个*  空格 2
 *******  4层7个*  空格 1
********* 5层9个*  空格 0
        规律：2n-1 规律：总层数-当前层数
*/
//4.将层数做成一个变量
//5.打印空心金字塔
/*
    *
   * *
  *   *
 *     *
*********
分析：在每行打印*是，需要考虑是打印*还是打印空格
结果：每层的第一个和最后一个是打印*，其他应该就是空的，即空格
最后一层例外，全部打*
*/
func main() {
	var totalLevel int = 9
	//i表示层数
	for i := 1; i <= totalLevel; i++ {
		//在打印*之间先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Printf(" ")
		}
		//j表示个数
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Printf("%s", "*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
