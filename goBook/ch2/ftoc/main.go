package main

import "fmt"

func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
func CToF(c float64) float64 {
	return c*9/5 + 32
}
func main() {
	var t float64
	fmt.Print("请输入你要转换的温度：")
	fmt.Scanln(&t)
	f := FToC(t)
	c := CToF(t)
	fmt.Printf("%.0f°C = %.2f°F, %.0f°C = %.2f°F\n", t, f, t, c)
}
