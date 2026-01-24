package main

import "fmt"

type Engine struct {
	NameEngine string
}

func (e Engine) Start() string {

	return fmt.Sprint("Завёлся:", e.NameEngine)
}

type Car struct {
	engine Engine
	Name   string
}

func (c Car) Drive() string {
	return c.engine.Start()
}

func main() {
	car := Car{engine: Engine{NameEngine: "JZ2"}, Name: "Toyota"}

	fmt.Println(car.Drive())
}
