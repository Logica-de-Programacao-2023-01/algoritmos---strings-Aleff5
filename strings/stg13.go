package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string

	fmt.Print("Digite uma sequencia numerica:")
	fmt.Scanln(&str)

	crescente := true
	for i := 1; i < len(str); i++ {
		numeroAnterior, err1 := strconv.Atoi(string(str[i-1]))
		numeroAtual, err2 := strconv.Atoi(string(str[i]))

		if err1 != nil || err2 != nil {
			fmt.Println("A sequência contém caracteres inválidos.")
			return
		}

		if numeroAtual <= numeroAnterior {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("A sequência é crescente.")
	} else {
		fmt.Println("A sequência não é crescente.")
	}
}
