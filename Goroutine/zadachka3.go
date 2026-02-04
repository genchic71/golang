package main

import (
	"fmt"
	"time"
)

func main() {
	//go func() {
	//	for i := range 3 {
	//		fmt.Println("Goroutine", i+1)
	//	}
	//}()
	//time.Sleep(100 * time.Millisecond)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Println("Goroutine", i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}
