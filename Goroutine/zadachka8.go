package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	wg.Wait()
}
