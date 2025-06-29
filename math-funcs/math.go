package main

import (
	"fmt"
	"strconv"
)

func operations(num1 int, num2 int) {
	fmt.Println(strconv.Itoa(num1) + " " + strconv.Itoa(num2))
	fmt.Println("Addition: ", num1+num2)
	fmt.Println("Subtraction: ", num1-num2)
	fmt.Println("Multiplication: ", num1*num2)
	fmt.Println("Division: ", num1/num2)
	fmt.Println("Remainder: ", num1%num2)
}

func main() {
	operations(10, 2)
}
