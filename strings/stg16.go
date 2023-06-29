package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Digite uma string:")
	fmt.Scanln(&str1)
	fmt.Print("Digite uma string:")
	fmt.Scanln(&str2)

	if strings.Contains(str1, str2) == true {
		fmt.Print("A segunda string é uma substring da primeira!")
	} else {
		fmt.Print("A segunda string não é uma substring da primeira!")
	}
}
