package main

import "fmt"

// interface本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (liHuaMao *Cat) Sleep() {
	fmt.Println("猫会睡觉")
}
func (liHuaMao *Cat) GetColor() string {
	return liHuaMao.color
}
func (liHuaMao *Cat) GetType() string {
	return "Cat"
}

// 具体类
type Dog struct {
	color string
}

func (haShiQi *Dog) Sleep() {
	fmt.Println("哈士奇会睡觉")
}
func (haShiQi *Dog) GetColor() string {
	return haShiQi.color
}
func (haShiQi *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("kind=", animal.GetType())
}

func main() {
	cat := Cat{color: "red"}
	showAnimal(&cat)
	dog := Dog{color: "yellow"}
	showAnimal(&dog)

}
