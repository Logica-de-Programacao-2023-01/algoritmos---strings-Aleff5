package main

import "fmt"

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	bytes := []byte(input)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	reversed := string(bytes)

	fmt.Println("String invertida:", reversed)
}
