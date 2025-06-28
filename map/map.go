package main

import "fmt"

var grades = make(map[string]int)

func addGrade(name string, grade int) {
	grades[name] = grade
}

func deleteGrade(name string) {
	delete(grades, name)
}

func getGrade(name string) int {
	return grades[name]
}

func main() {
	addGrade("Mark", 85)
	addGrade("Violet", 92)
	addGrade("Caitlyn", 78)

	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	viGrade := getGrade("Violet")
	fmt.Printf("\nVi's grade: %d\n", viGrade)

	deleteGrade("Caitlyn")
	fmt.Println("\nMap after deleting Caitlyn:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
