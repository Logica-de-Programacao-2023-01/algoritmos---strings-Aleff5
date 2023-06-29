package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Print("Digite uma string:")
	fmt.Scanln(&str)

	chars := []rune(str)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversedStr := string(chars)

	fmt.Print("A string ao contrario fica assim:", reversedStr)
}
