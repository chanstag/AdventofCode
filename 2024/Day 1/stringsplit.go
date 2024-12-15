package main

import (
	"fmt"
	"strings"
)

func main() {
	// Basic string splitting
	text := "Hello world programming"
	words := strings.Split(text, " ")
	fmt.Println(words) // Outputs: [Hello world programming]

	// Splitting with multiple spaces
	messyText := "  Hello   world  programming  "

	// Split with multiple consecutive spaces
	dirtyWords := strings.Split(messyText, " ")
	fmt.Println(dirtyWords) // Will have empty strings between multiple spaces

	// Clean splitting (remove empty strings)
	cleanWords := strings.Fields(messyText)
	fmt.Println(cleanWords) // Outputs: [Hello world programming]
}
