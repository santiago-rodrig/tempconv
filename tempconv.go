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
