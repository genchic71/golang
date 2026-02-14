package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ch <- 1

			fmt.Printf("Горутина %d: работает, занято слотов: %d\n",
				id, len(ch))
			time.Sleep(1 * time.Millisecond)
			<-ch
			fmt.Printf("Горутина %d: завершила работу, активных: %d\n", id, len(ch))
		}(i)
	}
	wg.Wait()

}
