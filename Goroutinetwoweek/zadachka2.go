package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Inc()
	Value() int
}

type SafeCounterstruct struct {
	count int
}

func (c *SafeCounterstruct) Inc(mutex *sync.Mutex) {
	mutex.Lock()
	c.count++
	mutex.Unlock()
}

func (c *SafeCounterstruct) Value(mutex *sync.Mutex) int {
	mutex.Lock()
	value := c.count
	mutex.Unlock()
	return value
}

func main() {
	mutex := &sync.Mutex{}
	counter := &SafeCounterstruct{count: 0}
	wg := &sync.WaitGroup{}
	wg.Add(50)
	for range 50 {
		go func() {
			defer wg.Done()
			counter.Inc(mutex)
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.Value(mutex))
}
