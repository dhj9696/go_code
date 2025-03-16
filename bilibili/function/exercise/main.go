package main

import "fmt"

func sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}
func swap(n1, n2 int) {
	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}
func swap2(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	fmt.Println("sum = ", sum(1, 1))
	swap(1, 2)

	a := 10
	b := 20
	swap2(&a, &b)
	fmt.Println(a, b)
}
