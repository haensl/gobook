package geometry

import "math"

// Point represents a two-dimensional point.
type Point struct{ X, Y float64 }

// Distance returns the distance between two points.
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance returns the distance between two points.
func (p Point) Distance(q Point) float64 {
	return Distance(p, q)
}
