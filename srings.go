package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is fast and go is easy"

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("Частота слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
