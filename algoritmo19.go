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
	j := len(x) - 1
	for i < j {
		y[i], y[j] = y[j], y[i]
		i++
		j--
	}
	fmt.Print("A palavra invertida é: ", string(y))
}
