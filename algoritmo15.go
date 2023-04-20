package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	i := 0
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&x)
	x = strings.ToLower(x)
	y := []byte(x)
	for i < len(y) {
		if strings.ContainsAny(string(y[i]), "aeiou") {
			y[i] = '*'
		}
		i++
	}
	fmt.Print(string(y))
}
