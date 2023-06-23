package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&input)

	wordCount := 0

	input = strings.TrimSpace(input)

	for i := 0; i < len(input); i++ {

		if input[i] == ' ' && i+1 < len(input) && input[i+1] != ' ' {
			wordCount++
		}
	}
	if len(input) > 0 {
		wordCount++
	}
	fmt.Printf("A frase contém %d palavras.\n", wordCount)
}
