// dup1 prints the count (in binary) and text of each line that
// appears more than once in the standard input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%b\t%s\n", n, line)
		}
	}
}
