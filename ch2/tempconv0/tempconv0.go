// Package tempconv encapsulates Celsius and Fahrenheit temperature conversions.
package tempconv

// Celsius degrees
type Celsius float64

// Fahrenheit degrees
type Fahrenheit float64

const (
	// AbsoluteZeroC represents the lowest degree Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC represents the freezing point of water in degrees Celsius
	FreezingC Celsius = 0
	// BoilingC represents the boiling point of water in degrees Celsius
	BoilingC Celsius = 100
)

// CToF converts degrees Celsius to degrees Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts degrees Fahrenheit to degrees Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
