package main

import "fmt"

func main() {
	var y []int
	for i := 0; i < 10; i++ {
		y = appendInt(y, i, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
