package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Я абсолютно счастлив, только не сегодня, а потом"
	splited := strings.Split(text, " ")

	for _, word := range splited {
		cleanedWord := strings.ReplaceAll(word, ",", "")
		fmt.Println(strings.ToLower(cleanedWord))
	}
}
