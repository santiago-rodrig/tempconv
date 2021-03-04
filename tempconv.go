// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius temperature scale
type Celsius float64

// Celsius.String returns a string representation of the Celsius type
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// Fahrenheit temperature scale
type Fahrenheit float64

// Fahrenheit.String returns a string representation of the Fahrenheit type
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// Kelvin temperature scale
type Kelvin float64

// Kelvin.String returns a string representation of the Kelvin type
func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

const (
	// AbsoluteZeroC is the absolute zero for the Celsius temperature scale
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing point for the Celsius temperature scale
	FreezingC Celsius = 0
	// BoilingC is the boiling point for the Celsius temperature scale
	BoilingC Celsius = 100
)

// CToF transforms Celsius temperatures into its Fahrenheit temperature equivalent
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC transforms Fahrenheit temperatures into its Celsius temperature equivalent
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK transforms Celsius temperatures into its Kelvin temperature equivalent
func CToK(c Celsius) Kelvin {
	return Kelvin(273.15 + c)
}

// KToC transforms Kelvin temperatures into its Celsius temperature equivalent
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF transforms Kelvin temperatures into its Fahrenheit temperature equivalent
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(KToC(k)*9/5 + 32)
}

// FToK transforms Fahrenheit temperatures into its Kelvin temperature equivalent
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(273.15 - FToC(f))
}
