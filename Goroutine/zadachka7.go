package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	go func() {
		ch <- "String"
		ch <- "Second"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
