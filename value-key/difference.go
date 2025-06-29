package main

import "fmt"

func changeByValue(first int, second int) {
	firstPtr := &first
	secondPtr := &second
	firstValue := *firstPtr
	first = *secondPtr
	second = firstValue
}

func changeByAddress(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}

func main() {
	a, b := 5, 3

	fmt.Println("Original values:")
	fmt.Printf("a = %d, b = %d\n\n", a, b)

	changeByValue(a, b)
	fmt.Println("After changeByValue:")
	fmt.Printf("a = %d, b = %d (no change)\n\n", a, b)

	changeByAddress(&a, &b)
	fmt.Println("After changeByAddress:")
	fmt.Printf("a = %d, b = %d (values swapped)\n", a, b)
}
