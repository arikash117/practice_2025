package main

import "fmt"

func multiply(num int) {
	fmt.Printf("\nMultiplication table for %d:\n", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

func main() {
	multiply(5)
}
