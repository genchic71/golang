package main

import "fmt"

// zad15
type List struct {
	len int
}

func (l *List) Len() int {
	if l == nil {
		return 0
	}
	return l.len
}

func main() {
	l := &List{len: 2}
	l = nil

	fmt.Println(l.Len())
}
