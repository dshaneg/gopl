package main

import (
	"fmt"
	"temperature/tempconv"
)

func main() {
	// f := tempconv.Fahrenheit(212.0)
	// c := tempconv.FToC(212.0)
	c := tempconv.Fahrenheit(212.0).ToCelcius()
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // no need to call String explicitly
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)   // does not call String
	fmt.Println(float64(c)) // does not call String
}
