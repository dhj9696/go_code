package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, val := range os.Args {
		if idx == 0 {
			fmt.Println("os.Args[0]=", val)
		}
		fmt.Println(idx, "=", val)
	}
}
