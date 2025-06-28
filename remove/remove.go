package main

import "fmt"

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	songs := []string{"Лучшее", "Эфир", "Альтернатива", "Впереди", "Крепость"}

	fmt.Println("Songs:", songs)

	songs = removeElement(songs, 2)
	fmt.Println("After removal:", songs)

	songs = removeElement(songs, 0)
	fmt.Println("Remove first element:", songs)

	songs = removeElement(songs, len(songs)-1)
	fmt.Println("Remove last element:", songs)
}
