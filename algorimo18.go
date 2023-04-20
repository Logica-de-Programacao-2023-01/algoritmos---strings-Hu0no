package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	fmt.Print("Insira uma string: ")
	fmt.Scan(&x)
	y, err := strconv.Atoi(x)
	if err == nil && y == y {
		fmt.Print("Há somente números na string")
	} else {
		fmt.Print("Não há somente números na string")
	}
}
