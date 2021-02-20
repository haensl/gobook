package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	ticker := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("%d\n", countdown)
		select {
		case <-ticker.C:
		case <-abort:
			fmt.Println("Launch aborted!")
			ticker.Stop()
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
