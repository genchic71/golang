package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Println("Goroutine", i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}
