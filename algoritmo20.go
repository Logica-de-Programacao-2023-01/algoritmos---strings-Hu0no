package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto: ")
	scanner.Scan()
	x := scanner.Text()
	y := strings.Split(x, " ")
	for _, word := range y {
		for i, char := range word {
			if i == 0 && !unicode.IsUpper(char) {
				fmt.Print("O texto não está em CamelCase")
				return
			} else if i != 0 && !unicode.IsLower(char) {
				fmt.Print("O texto não está em CamelCase")
				return
			}
		}
	}
	fmt.Print("O texto está em CamelCase")
}
