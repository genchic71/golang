package main

import (
	"fmt"
	"sync"
)

func main() {
	var m int
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch <- i
		}
	}()
	go func(m *int) {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			*m += <-ch
		}
	}(&m)
	wg.Wait()
	fmt.Println(m)
}
