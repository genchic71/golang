package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine", i)
		}(i)
	}
	wg.Wait()
}
