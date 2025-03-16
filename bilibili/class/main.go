package main

import "fmt"

type Hero struct {
	Name  string
	Age   int
	Level int
}

//	func (person Hero) Show() {
//		fmt.Printf("name=%s, age=%d, level=%d\n", person.Name, person.Age, person.Level)
//	}
//
//	func (person Hero) GetName() string {
//		return person.Name
//	}
//
//	func (person Hero) SetName(newName string) {
//		person.Name = newName
//	}
func (person *Hero) Show() {
	fmt.Printf("name=%s, age=%d, level=%d\n", person.Name, person.Age, person.Level)
}
func (person *Hero) GetName() string {
	return person.Name
}
func (person *Hero) SetName(newName string) {
	person.Name = newName
}
func main() {
	//创建一个对象
	person := Hero{Name: "贾登恒", Age: 18, Level: 10}
	person.Show()

	person.SetName("李四") //并没有修改成功，需要使用指针的形式才能修改成功
	person.Show()

}
