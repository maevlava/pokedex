package utils

import "strings"

func CleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Split(trimmedText, " ")

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}
