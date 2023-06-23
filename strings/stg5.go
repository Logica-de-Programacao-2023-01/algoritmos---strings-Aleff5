package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&input)

	_, err := strconv.ParseFloat(input, 64)
	if err == nil {
		fmt.Println("A string é um número válido em ponto flutuante.")
	} else {
		fmt.Println("A string não é um número válido em ponto flutuante.")
	}
}
