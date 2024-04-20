package main

import (
	"fmt"
	"temperature/tempconv"
)

func main() {
	c := tempconv.Fahrenheit(212.0).ToCelsius()
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // no need to call String explicitly
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)   // does not call String
	fmt.Println(float64(c)) // does not call String
	fmt.Println(c.ToKelvin())
	fmt.Println(tempconv.Kelvin(0).ToCelsius())
	fmt.Println(tempconv.Kelvin(0).ToFahrenheit())
}
