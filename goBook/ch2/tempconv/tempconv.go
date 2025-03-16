// Package tempconv performs Celsius and Fahrenheit conversions.
// 把变量的声明、对应的常量，还有方法都放到tempconv.go源文件中
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroCelsius Celsius = -273.15
	FreezingCelsius     Celsius = 0
	BoilingCelsius      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
