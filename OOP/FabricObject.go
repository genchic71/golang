package main

import (
	"fmt"
)

// zad 11
type Person struct {
	Name string
	Age  int
}

func Newperson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	person := Newperson("John", 30)
	fmt.Println(person)
}
