package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var veioChar, nvChar byte

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanf("%c\n", &veioChar)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanf("%c\n", &nvChar)

	output := strings.ReplaceAll(input, string(veioChar), string(nvChar))

	fmt.Println("Resultado:", output)
}
