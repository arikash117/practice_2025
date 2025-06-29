package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findElement(slice []string, el string) int {
	for i, e := range slice {
		if e == el {
			return i
		}
	}
	return -1
}

func sortSlice(slice []string) {
	sort.Strings(slice)
}

func filterSlice(slice []string, condition func(string) bool) []string {
	var result []string
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	words := []string{"banana", "apple", "cherry", "grape"}

	index := findElement(words, "apple")
	fmt.Println("find: " + strconv.Itoa(index))

	sortSlice(words)
	fmt.Println("Sorted: ", words)

	longWords := filterSlice(words, func(s string) bool {
		return len(s) > 5
	})
	fmt.Println("filtered: ", longWords)
}
