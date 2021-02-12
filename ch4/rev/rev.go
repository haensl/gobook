// Package reverse contains slice reveral functionality
package reverse

// Reverse reverses a slice of ints in place
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ShiftLeft rotates s left by n positions
func ShiftLeft(s []int, n int) {
	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s)
}
