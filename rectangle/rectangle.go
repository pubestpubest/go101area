package rectangle

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width  float32
	height float32
}

func NewRectangle(w, h float32) (Rectangle, error) {
	if w < 0 || h < 0 {
		return Rectangle{}, errors.New("side(s) must be positive")
	} else {
		return Rectangle{w, h}, nil
	}
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
