package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 2; i < 101; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
