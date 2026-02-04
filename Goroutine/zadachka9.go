package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	stop := false

	go func() {
		<-done
		stop = true
	}()

	go func(done chan bool) {
		for !stop {
			fmt.Println("Tick")
			time.Sleep(300 * time.Millisecond)
		}
	}(done)

	time.Sleep(1500 * time.Millisecond)
	done <- true
	fmt.Println("Stopped")
}
