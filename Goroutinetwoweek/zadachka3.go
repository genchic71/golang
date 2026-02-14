package main

import (
	"fmt"
	"sync"
)

type SafeCounter interface {
	Inc()
	Value() int
}

type Counter struct {
	count int
}

func (c *Counter) Inc(mutex *sync.RWMutex) {
	mutex.Lock()
	c.count++
	mutex.Unlock()
}

func (c *Counter) Value(mutex *sync.RWMutex) int {
	mutex.RLock()
	value := c.count
	mutex.RUnlock()
	return value
}

func main() {
	mutex := &sync.RWMutex{}
	counter := &Counter{count: 0}
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
