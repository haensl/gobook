package bytecounter

// ByteCounter counts written bytes
type ByteCounter int

// Write writes p bytes to the ByteCounter
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
