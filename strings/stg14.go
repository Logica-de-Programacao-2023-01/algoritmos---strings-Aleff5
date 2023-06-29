package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string

	fmt.Print("Digite uma sequencia numerica:")
	fmt.Scanln(&str)

	decrescente := true
	for i := 1; i < len(str); i++ {
		numeroAnterior, err1 := strconv.Atoi(string(str[i-1]))
		numeroAtual, err2 := strconv.Atoi(string(str[i]))

		if err1 != nil || err2 != nil {
			fmt.Println("A sequência contém caracteres inválidos.")
			return
		}

		if numeroAnterior <= numeroAtual {
			decrescente = false
			break
		}
	}

	if decrescente {
		fmt.Println("A sequência é decrescente.")
	} else {
		fmt.Println("A sequência não é decrescente.")
	}
}
