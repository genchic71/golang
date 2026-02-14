package main

import (
	"fmt"
	"log"
	"sync"
)

type SafeLogger struct {
	sync.Mutex
}

func (sl *SafeLogger) Log(msg string) {
	sl.Lock()
	defer sl.Unlock()
	log.Println(msg)
}

func main() {
	logger := &SafeLogger{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				logger.Log(fmt.Sprintf("%d %d", id, j))
			}
		}(i)
	}
	wg.Wait()
}
