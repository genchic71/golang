package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	done := make(chan bool, 1)
	//go func() {
	//	ch <- "Worker done"
	//}()
	//v, ok := <-ch
	//if !ok {
	//	return
	//}
	//fmt.Println(v)
	//fmt.Println("Main continues")
	go func(done chan bool) {
		ch <- "Worker done"
		done <- true
	}(done)
	<-done
	v := <-ch
	fmt.Println(v)
	fmt.Println("Main continues")
}
