package square

import (
	"errors"
	"fmt"
)

type Square struct {
	width float32
}

func NewSquare(side float32) (Square, error) {
	if side < 0 {
		return Square{}, errors.New("side must be positive")
	} else {
		return Square{side}, nil
	}
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
