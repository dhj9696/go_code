package main

import (
	"fmt"
	"go_code/ch10/user"
) //引用包的路径

func main() {
	c := user.Course{
		Name: "golang",
	}
	fmt.Println(user.GetCourse(c))
}
