package main

import "fmt"

func main() {
	s1 := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	s2 := fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	fmt.Println(s3)
	fmt.Printf("%d", "hi")
}
