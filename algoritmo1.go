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
	x := scanner.Text()
	y := strings.ToUpper(x)
	fmt.Print("Sua frase em letras maiúsculas é: ", y)
}
