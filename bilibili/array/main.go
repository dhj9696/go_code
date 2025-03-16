package main

import "fmt"

/*案例1：创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。使用for循环访问所有
  元素并打印出来。提示：字符串数据运算：'A'+ -> 'B'
*/
//思路
//1.声明一个数组var myChars [26]byte
//2.使用for循环，利用字符可以进行运算的特点来赋值'A'+1 -> 'B'
//3.使用for循环打印即可
func main() {
	//var myChars [26]byte
	//for i := 0; i < 26; i++ {
	//	myChars[i] = byte('A' + i)
	//	fmt.Printf("%c ", myChars[i])
	//}
	//求出一个数组的最大值，并得到对应的下标
	//思路
	//1.声明一个数组var intArr[5] = [...]int {1, -1, 9, 90, 11}
	//2.假设第一个元素就是最大值，下标就是0
	//3.然后从第二个元素开始循环比较，如果发现有更大，则交换
	var intArr = [...]int{1, -5, 39, 100, 90}
	maxValue := intArr[0]
	maxIndex := 0
	for i := 1; i < len(intArr); i++ {
		if maxValue < intArr[i] {
			maxValue = intArr[i]
			maxIndex = i
		}
	}
	fmt.Println(maxValue, maxIndex)

}
