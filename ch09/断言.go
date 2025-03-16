package main

import "fmt"

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai32, _ := a.(int32)
		bi32, _ := b.(int32)
		return ai32 + bi32
	case int64:
		ai64, _ := a.(int64)
		bi64, _ := b.(int64)
		return ai64 + bi64
	case float32:
		af32, _ := a.(float32)
		bf32, _ := b.(float32)
		return af32 + bf32
	case float64:
		af64, _ := a.(float64)
		bf64, _ := b.(float64)
		return af64 + bf64
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("not supported type")
	}
}
func main() {
	a := 10.1
	b := 20.3
	c := add(a, b)
	fmt.Println(c)
}
