package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		digest := sha256.Sum256([]byte(arg))
		fmt.Printf("%x\n%[1]T\n", digest)
	}
}
