package main

import (
	"fmt"
	"log"

	"github.com/pubestpubest/go101area/circle"
	"github.com/pubestpubest/go101area/rectangle"
	"github.com/pubestpubest/go101area/square"
)

type Shape interface {
	Area() (float32, error)
	Introduce() string
}

func handleError[T Shape](value T, er error) T {
	if er != nil {
		log.Fatalf("Error: %v", er)
	}
	return value
}

func main() {
	square := handleError(square.NewSquare(3))
	rec := handleError(rectangle.NewRectangle(-2, 2))
	circle := handleError(circle.NewCircle(1))
	shapes := []Shape{circle, rec, square}
	for _, shape := range shapes {
		area, error := shape.Area()
		if error != nil {
			fmt.Println("Error: ", error)
			return
		}
		fmt.Printf("%-40s area: %.2f\n", shape.Introduce(), area)
	}
}
