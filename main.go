package main

import (
	"errors"
	"fmt"
)

type Area interface {
	Area() (float32, error)
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

func (c Circle) Area() (float32, error) {
	if c.radius < 0 {
		return 0, errors.New("radius must be positive")
	}
	return PI * c.radius * c.radius, nil
}

func (c Circle) Introduce() string {
	return fmt.Sprintf("Circle with radius: %.2f", c.radius)
}

func (r Rectangle) Area() (float32, error) {
	if r.height < 0 || r.width < 0 {
		return 0, errors.New("side number must be positive")
	}
	return r.width * r.height, nil
}

func (r Rectangle) Introduce() string {
	return fmt.Sprintf("Rectangle with width: %.2f, height: %.2f", r.width, r.height)
}

func (s Square) Area() (float32, error) {
	if s.width < 0 {
		return 0, errors.New("side number must be positive")
	}
	return s.width * s.width, nil
}

func (s Square) Introduce() string {
	return fmt.Sprintf("Square with width: %.2f", s.width)
}

func main() {
	squre := Square{3}
	rec := Rectangle{-2, 2}
	circle := Circle{1}
	shapes := []Area{circle, rec, squre}
	for _, shape := range shapes {
		area, error := shape.Area()
		if error != nil {
			fmt.Println("Error: ", error)
			return
		}
		fmt.Printf("%-40s area: %.2f\n", shape.Introduce(), area)
	}
}
