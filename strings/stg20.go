package main

import (
	"fmt"
	"unicode"
)

func main() {
	var entrada string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&entrada)

	isCamelCase, numPalavras := verificaCamelCase(entrada)

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", numPalavras)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func verificaCamelCase(str string) (bool, int) {
	if len(str) > 0 && !unicode.IsUpper(rune(str[0])) {
		return false, 0
	}
	numPalavras := 1

	for i := 1; i < len(str); i++ {
		if unicode.IsUpper(rune(str[i])) {
			numPalavras++
		}
	}

	return true, numPalavras
}
