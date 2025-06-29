package main

import "fmt"

func weekDay(num int) {
	switch num {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("?")
	}
}

func main() {
	var day int
	fmt.Print("enter day number (1-7): ")
	fmt.Scan(&day)
	weekDay(day)
}
