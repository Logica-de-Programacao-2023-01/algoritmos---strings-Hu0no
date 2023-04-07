package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto: ")
	scanner.Scan()
	texto := scanner.Text()
	newtexto := strings.ReplaceAll(texto, " ", "")
	fmt.Print("O texto sem espaços é: ", newtexto)
}
