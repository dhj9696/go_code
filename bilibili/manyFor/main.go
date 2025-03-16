package main

import "fmt"

/*
编程两大绝招：
(1) 先易后难，即将一个复杂的问题分解成简单的问题。
(2) 先死后活，即将代码先用最简单的办法实现功能，后在使用函数变量等将代码活起来
*/

//多重循环控制
/*
案例一：统计3个班成绩情况，每个班有5名同学，求出各个班的平均分和所有
班级的平均分[学生的成绩从键盘输入]
*/
//分析实现思路1
//1.统计1个班级成绩情况，每个班有5名同学，求出该班的平均分[学生成绩由键盘输入]
//2.学生数就是5个[先死后活]
//3.声明一个sum统计班级的总分
//分析实现思路2
//1.统计3个班级成绩情况，每个班有5名同学，求出该班的平均分[学生成绩由键盘输入]
//2.j表示第几个班级
//3.定义一个变量存放总成绩

func main() {
	var sumAll float64
	var classNumbers int = 3
	var stuNumbers int = 5
	var passCount int = 0
	for j := 1; j <= classNumbers; j++ {
		sum := 0.0
		for i := 1; i <= stuNumbers; i++ {
			var score float64
			fmt.Printf("请输入第%d个 班第%d个学生的成绩: \n", j, i)
			_, err := fmt.Scanf("%f", &score)
			if err != nil {
				fmt.Println("输入无效，请输入一个有效的数字。")
			}
			if score >= 60 {
				passCount++
			}
			//累计总分
			sum += score
		}
		sumAll += sum
		fmt.Printf("第%d个班级的平均分是: %v\n", j, sum/float64(classNumbers))
	}
	fmt.Printf("所有班的总成绩为: %v \n", sumAll)
	fmt.Printf("所有班的平均成绩为: %v\n", sumAll/float64(classNumbers*stuNumbers))
	fmt.Printf("所有班及格人数为: %d\n", passCount)
}
