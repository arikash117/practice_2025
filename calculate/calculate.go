package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Scan(&num2)

	var result float64
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err = fmt.Errorf("division by zero is impossible")
		} else {
			result = num1 / num2
		}
	default:
		err = fmt.Errorf("unknown operator: %s", operator)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
