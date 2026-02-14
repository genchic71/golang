package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Stats struct {
	requests int
	errors   int
}

func (s *Stats) RecordRequest(r int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	s.requests += r

}

func (s *Stats) RecordError(e int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	s.errors += e

}

func main() {
	mutex := &sync.Mutex{}
	s := Stats{requests: 0, errors: 0}
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			r := rand.Intn(20)
			defer wg.Done()
			if r < 18 {
				s.RecordRequest(1, mutex)
			} else {
				s.RecordError(1, mutex)
			}
		}()
	}
	wg.Wait()
	fmt.Println(s)
}
