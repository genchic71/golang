package main

import (
	"fmt"
	"sync"
)

var Counter int

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			mutex.Lock()
			Counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", Counter)
}
