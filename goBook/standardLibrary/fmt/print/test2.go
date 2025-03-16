package main

import (
	"fmt"
	"math"
)

func main() {
	integer := 23
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	fmt.Printf("%T %T\n", integer, &integer)

	truch := true
	fmt.Printf("%v %t\n", truch, truch)

	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)

	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)

	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
}
