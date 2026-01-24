package main

import "fmt"

type dsjhvc struct {
	name string
}

func (d dsjhvc) Speak() string {
	return fmt.Sprint("Woof")
}

type cdgcsgh struct {
	name2 string
}

func (d cdgcsgh) Speak() string {
	return fmt.Sprint("Гав")
}

type fbsjfgvj struct {
	cdgcsgh
	dsjhvc
}

func main() {
	fsad := fbsjfgvj{cdgcsgh{name2: "hjfv"}, dsjhvc{name: "hjdv"}}
	fmt.Println(fsad.dsjhvc.Speak())
}
