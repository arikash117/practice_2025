package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	Age    int
	Course int
	Grades []int
}

func NewStudent(name string, age, course int, grades []int) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
		Grades: grades,
	}
}

func (student Student) getName() string {
	return student.Name
}

func (student Student) getAge() int {
	return student.Age
}

func (student Student) getAverageGrade() float64 {
	if len(student.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range student.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(student.Grades))
}

func (student Student) GetStatus() string {
	average := student.getAverageGrade()
	switch {
	case average >= 4.5:
		return "piltovers finest nerd"
	case average >= 3.5:
		return "piltovers finest"
	default:
		return "okk..."
	}
}

func main() {
	student1 := NewStudent("Caitlyn Kiramman", 20, 2, []int{5, 4, 5, 5, 5, 5, 5})
	student2 := NewStudent("Violet Kiramman", 20, 1, []int{4, 3, 2, 4, 5, 4, 5})

	fmt.Println("Student " + student1.getName() + " is " + student1.GetStatus() + " and " + strconv.Itoa(student1.getAge()) + "y.o.")
	fmt.Println("Student " + student2.getName() + " is " + student2.GetStatus() + " and " + strconv.Itoa(student2.getAge()) + "y.o.")

}
