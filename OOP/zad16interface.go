package main

import "fmt"

type chiptuning interface {
	Power()
}

type skoda struct {
	power int
}

func (s *skoda) Power() {
	s.power += 44
}

func main() {
	s := &skoda{power: 154}
	var c chiptuning = s
	c.Power()
	fmt.Println(s)

}
