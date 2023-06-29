package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string:")
	fmt.Scanln(&str)

	unicas := ""
	for _, char := range str {
		count := strings.Count(str, string(char))
		if count == 1 {
			unicas += string(char)
		}
	}
	fmt.Print("As letras que são unicas na string inserida são:", unicas)
}
