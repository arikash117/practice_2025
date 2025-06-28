package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mark string = "Hoegaarden"
	var liters float32 = 0.5
	var amount int = 2
	var isEnough bool
	if liters*2 >= 1 {
		isEnough = true
	}
	if isEnough {
		fmt.Println("Let's go to a walk and bring " + strconv.Itoa(amount) + " bottles of " + mark)
	}
}
