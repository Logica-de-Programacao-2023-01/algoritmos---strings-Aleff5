package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&entrada)

	withoutVowels := removeVowels(entrada)
	fmt.Println("String sem vogais:", withoutVowels)
}

func removeVowels(str string) string {

	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, vowel, "")
	}

	return str
}
