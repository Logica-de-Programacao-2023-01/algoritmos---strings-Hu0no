package main

import "fmt"

func main() {
	var x string
	fmt.Print("Insira uma string: ")
	fmt.Scan(&x)
	y := 0
	z := 1
	for z < len(x) {
		if x[y] > x[z] {
			fmt.Print("A string não é uma sequência numérica crescente")
			return
		}
		y++
		z++
	}
	fmt.Print("A string é uma sequência numérica crescente")
}
