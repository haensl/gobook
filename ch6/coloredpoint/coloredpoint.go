package coloredpoint

import (
	"haensl/gobook/ch6/geometry"
	"image/color"
)

// ColoredPoint represents a two-dimensional point with a color
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
