package circle

import (
	"errors"
	"fmt"
)

type Circle struct {
	radius float32
}

const PI = 3.14

func NewCircle(r float32) (Circle, error) {
	if r < 0 {
		return Circle{}, errors.New("radius must be positive")
	} else {
		return Circle{r}, nil
	}
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
