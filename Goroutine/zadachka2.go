package main

import (
	"fmt"
	"time"
)

func greet(a string) {
	fmt.Println("Hello", a)
}

func main() {
	a := "Alice"
	go greet(a)
	time.Sleep(100 * time.Millisecond)
}
