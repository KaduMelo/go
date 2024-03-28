package main

import "fmt"

type Celsius float64
type Fahrenheit float32

const (
	AbsoluteC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) error compilation
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
