package main

import "fmt"

func main() {
	ch := make(chan int)
	var v int

	go func() {
		ch <- 42
	}()

	v = <-ch
	fmt.Println("Received:", v)
}
