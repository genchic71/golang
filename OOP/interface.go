package main

import "fmt"

// полностью 3 уровень тут
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}
type Circle struct {
	Radius float64
}
type Square struct {
	X, Y float64
}

func (s Square) Area() float64 {
	return s.X*s.X + s.Y*s.Y
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r *Circle) Area() float64 {
	return r.Radius * r.Radius
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func Describe(v interface{}) {
	switch i := v.(type) {
	case int:
		fmt.Println("передали int")
	case string:
		fmt.Println("передали string")
	case float64:
		fmt.Println("передали float64")
	case *Rectangle:
		fmt.Println("передали Rectangle")
	default:
		fmt.Printf("передали неизвестный тип: %T\n", i)
	}
}

func What(s interface{}) {
	s, ok := s.(Square)
	if ok {
		fmt.Println("преобраховали Square")
	} else {
		fmt.Println("Не очень")
	}
}

func main() {
	var i interface{}
	var b interface{}
	r := &Rectangle{
		Width:  3,
		Height: 4,
	}
	s := Square{X: 5, Y: 5}
	PrintArea(r)
	i = s
	b = r
	What(i)
	Describe(b)
}
