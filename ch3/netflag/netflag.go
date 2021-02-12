// Package netflag contains network interface flags and manipulations
package netflag

// Flags represent network interface flags
type Flags uint

const (
	// FlagUp interface is up
	FlagUp Flags = 1 << iota
	// FlagBroadcast interface is broadcasting
	FlagBroadcast
	// FlagLoopback interface is loopback
	FlagLoopback
	// FlagPointToPoint interface is point to point
	FlagPointToPoint
	// FlagMulticast interface is multicast
	FlagMulticast
)

// IsUp checks whether an interface is up
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// TurnDown disables an interface
func TurnDown(v *Flags) { *v &^= FlagUp }

// SetBroadcast turns an interface into a broadcast interface
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

// IsCast determines whether or not an interface is broadcasting or multicasting
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }
