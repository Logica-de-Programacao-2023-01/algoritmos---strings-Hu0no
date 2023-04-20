package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	var y string
	fmt.Print("Insira uma palavra: ")
	fmt.Scan(&x)
	fmt.Print("Insira uma letra: ")
	fmt.Scan(&y)
	if strings.Contains(x, y) {
		fmt.Println("Índices que a letra aparece: ")
		for i := range x {
			if string(x[i]) == y {
				fmt.Println(i)
			}
		}
		return
	}
	fmt.Print("A letra que você digitou não existe nessa palavra")
}
