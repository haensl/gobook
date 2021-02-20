package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	var timestamp time.Time
	for countdown := 10; countdown > 0; countdown-- {
		timestamp = <-tick
		fmt.Printf("%d: %v\n", countdown, timestamp)
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
