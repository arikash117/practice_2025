package main

import "fmt"

func main() {
	var songs []string

	songs = append(songs, "Лучшее")
	songs = append(songs, "Эфир")
	songs = append(songs, "Альтернатива")

	fmt.Println(songs)
	fmt.Println("Lenght:", len(songs))
	fmt.Println("Capacity:", cap(songs))

	fmt.Println("\nSlice elements:")
	for _, song := range songs {
		fmt.Println(song)
	}
}
