package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	//第一种嵌套方式
	//person Person
	//score  float32
	//第二种嵌套方式，匿名嵌套
	Person
	score float32
	name  string
}

func main() {
	//s := Student{
	//	person: Person{
	//		"jdh", 18,
	//	},
	//	score: 98.4,
	//}
	s := Student{
		Person{
			"jdh", 21,
		},
		96.4,
		"jdh2",
	}
	s.name = "jdh3"
	fmt.Println(s)
	fmt.Println(s.name)
}
