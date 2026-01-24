package main

import "fmt"

type Animel struct {
	name string
}

func (a Animel) Speak() string {
	return fmt.Sprint(a.name)
}

type Dog struct {
	Animel
	Age int
}

func (d Dog) Speak() string {
	return fmt.Sprint("Woof")
}

func main() {
	Dog := Dog{Animel: Animel{name: "Bod"}, Age: 10}
	animel := Animel{name: "JOB"}
	fmt.Println(animel.Speak())
	fmt.Println(Dog.Speak())
}
