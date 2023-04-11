package main

import "fmt"

func main() {
	var x string
	fmt.Print("Insira uma palavra: ")
	fmt.Scan(&x)
	newx := ""
	for i := 0; i < len(x); i++ {
		if string(x[i]) != "a" && string(x[i]) != "e" && string(x[i]) != "i" && string(x[i]) != "o" && string(x[i]) != "u" {
			newx += string(x[i])
		}
	}
	fmt.Print("A palavra sem vogais é: ", newx)
}
