package main

import (
	"fmt"
	"math"
)

type Form interface {
	Area() float32
}

type Square struct {
	height int
	width  int
}

type Circle struct {
	radius int
}

func (s Square) Area() float32 {
	return float32(s.height * s.width)
}

func (c Circle) Area() float32 {
	return float32(c.radius*c.radius) * math.Pi
}

func main() {
	square := Square{height: 5, width: 5}
	circle := Circle{radius: 5}

	forms := []Form{square, circle}

	for _, form := range forms {
		fmt.Printf("Form area : %.2f\n", form.Area())
	}
}
