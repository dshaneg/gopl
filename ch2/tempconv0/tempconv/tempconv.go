// Package tempconv performs Celsius, Kelvin, and Fahrenheit temperature conversions
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (k Kelvin) ToCelsius() Celsius       { return Celsius(k) + AbsoluteZeroC }
func (k Kelvin) ToFahrenheit() Fahrenheit { return k.ToCelsius().ToFahrenheit() }
func (k Kelvin) String() string           { return fmt.Sprintf("%gK", k) }

func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) ToKelvin() Kelvin         { return Kelvin(c - AbsoluteZeroC) }
func (c Celsius) String() string           { return fmt.Sprintf("%g℃", c) }

func (f Fahrenheit) ToCelsius() Celsius { return Celsius((f - 32) * 5 / 9) }
func (f Fahrenheit) ToKelvin() Kelvin   { return f.ToCelsius().ToKelvin() }
func (f Fahrenheit) String() string     { return fmt.Sprintf("%g℉", f) }
