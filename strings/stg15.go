package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string:")
	fmt.Scanln(&str)

	str = strings.Replace(str, "A", "*", -1)
	str = strings.Replace(str, "a", "*", -1)
	str = strings.Replace(str, "E", "*", -1)
	str = strings.Replace(str, "e", "*", -1)
	str = strings.Replace(str, "I", "*", -1)
	str = strings.Replace(str, "i", "*", -1)
	str = strings.Replace(str, "O", "*", -1)
	str = strings.Replace(str, "o", "*", -1)
	str = strings.Replace(str, "U", "*", -1)
	str = strings.Replace(str, "u", "*", -1)

	fmt.Println("A nova string é:", str)
}
