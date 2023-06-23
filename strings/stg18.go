package main

import "fmt"

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if contemApenasNumeros(input) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}
func contemApenasNumeros(str string) bool {

	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
