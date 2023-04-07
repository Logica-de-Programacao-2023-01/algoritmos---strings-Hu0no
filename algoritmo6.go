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
	t := scanner.Text()
	fmt.Print(strings.Count(t, " ")+1, " é o número de palavras. ")
}
