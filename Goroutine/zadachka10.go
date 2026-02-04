package main

import (
	"fmt"
	"sync"
)

func Zapice(wg *sync.WaitGroup) chan int {
	ch := make(chan int, 3)
	go func() {
		defer wg.Done()
		for i := range 3 {
			ch <- i + 1
		}
		close(ch)
	}()
	return ch
}

func Poluch(ch chan int, wg *sync.WaitGroup) {

	go func() {
		defer wg.Done()
		for range 3 {
			v := <-ch
			fmt.Println(v)
		}
	}()

}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	Poluch(Zapice(wg), wg)
	wg.Wait()
}
