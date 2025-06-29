package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	Age    int
	Course int
}

func NewStudent(name string, age, course int) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
	}
}

func (student Student) getInfo() string {
	return "name: " + student.Name + ", age: " + strconv.Itoa(student.Age)
}

func (student Student) upgrade() {
	student.Course++
}

func main() {
	student1 := NewStudent("Caitlyn Kiramman", 20, 2)
	student2 := NewStudent("Violet Kiramman", 20, 1)
	student2.upgrade()
	fmt.Println(student1.getInfo() + " and " + student2.getInfo() + " have the same last name because they are married <3")
}
