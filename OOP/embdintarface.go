package main

import "fmt"

// zad18
type Reader interface {
	Read()
}

type Closer interface {
	Close()
}
type ReadCloser interface {
	Reader
	Closer
}

type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println(f.Name)
}
func (f File) Close() {
	fmt.Println("закрыто")
}

func reclo(r ReadCloser) {
	r.Read()
	r.Close()
}
func main() {
	f := File{Name: "100 ошибок go"}
	reclo(f)
}
