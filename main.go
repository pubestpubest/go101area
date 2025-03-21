package main

import (
	"fmt"
)

type Area interface {
	Area() float32
	Introduce() string
}

const PI = 3.14

type Circle struct {
	radius float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Square struct {
	width float32
}

func (c Circle) Area() float32 {
	return PI * c.radius * c.radius
}

func (c Circle) Introduce() string {
	return fmt.Sprintf("Circle with radius: %.2f", c.radius)
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func (r Rectangle) Introduce() string {
	return fmt.Sprintf("Rectangle with width: %.2f, height: %.2f", r.width, r.height)
}

func (s Square) Area() float32 {
	return s.width * s.width
}

func (s Square) Introduce() string {
	return fmt.Sprintf("Square with width: %.2f", s.width)
}

func main() {
	squre := Square{3}
	rec := Rectangle{2, 2}
	circle := Circle{1}
	shapes := []Area{circle, rec, squre}
	for _, shape := range shapes {
		fmt.Printf("%-40s area: %.2f\n", shape.Introduce(), shape.Area())
	}
}
