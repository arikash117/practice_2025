package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Я абсолютно счастлив, только не сегодня, а потом"
	str := "счаст"

	fmt.Printf("Length: %d\n", len(text))
	fmt.Printf("Contains a substring \"%s\": %t\n", str, strings.Contains(text, str))
	fmt.Printf("Upper case: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower case: %s\n", strings.ToLower(text))
}
