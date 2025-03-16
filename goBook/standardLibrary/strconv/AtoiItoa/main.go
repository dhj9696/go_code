package main

import (
	"fmt"
	"strconv"
)

//Atoi(string to int) Itoa(int to string)

func main() {
	i, err := strconv.Atoi("-42")
	if err != nil {
		panic(err)
	}
	s := strconv.Itoa(-42)
	fmt.Printf("s = %q\n", s) //"42"
	fmt.Printf("i = %d\n", i) //42
}
