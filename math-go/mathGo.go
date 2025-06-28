package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int = 4
	var circleArea float32 = float32(radius*2) * math.Pi
	fmt.Printf("Area of a circle = %.2f\n", circleArea)

	const e = math.E
	naturalLog := math.Log(e)
	fmt.Printf("ln(e) = %.2f\n", naturalLog)
}
