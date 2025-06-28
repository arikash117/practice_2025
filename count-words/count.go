package main

import (
	"fmt"
	"strings"
)

var wordCount = make(map[string]int)
var text = "бла бла бла бле бле бле бли бли бли линган гули гули гули гваджа линган гу линган гу линган гули гули гули гваджа линган гу линган гу"

func count(text string, wordCountList map[string]int) map[string]int {
	for _, word := range strings.Split(text, " ") {
		wordCountList[word] = strings.Count(text, word)
	}
	return wordCountList
}

func main() {
	count(text, wordCount)
	fmt.Println(wordCount)
}
