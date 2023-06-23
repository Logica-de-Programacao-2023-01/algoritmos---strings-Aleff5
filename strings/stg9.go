package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada, veiochar, nvchar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&entrada)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&veiochar)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&nvchar)

	// Substituir todas as ocorrências da letra na string
	replaced := strings.ReplaceAll(entrada, veiochar, nvchar)

	// Exibir o resultado ao usuário
	fmt.Println("String com substituição:", replaced)
}
