package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Main started")

	go func() {
		fmt.Println("Goroutine finished")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main finished")
}
