package main

import "fmt"

func change(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}

func main() {

	a := 10
	b := 4

	fmt.Printf("before change: a = %d, b = %d\n", a, b)

	change(&a, &b)

	fmt.Printf("after change: a = %d, b = %d\n", a, b)
}
