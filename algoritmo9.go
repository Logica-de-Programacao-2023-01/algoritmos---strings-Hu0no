package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var l string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto: ")
	scanner.Scan()
	texto := scanner.Text()
	fmt.Print("Escolha uma letra para ser removida do texto: ")
	fmt.Scan(&l)
	newtexto := strings.ReplaceAll(texto, l, "")
	fmt.Print("O texto com as modificações é: ", newtexto)
}
