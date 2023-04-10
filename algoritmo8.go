package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto: ")
	scanner.Scan()
	str := scanner.Text()
	strRev := reverse(str)
	fmt.Print("O texto invertido é: ", strRev)
}
