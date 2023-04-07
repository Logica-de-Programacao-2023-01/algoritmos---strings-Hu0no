package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto ou um número: ")
	scanner.Scan()
	x := scanner.Text()
	i, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Printf("O valor %s não é um float\n", x)
	} else {
		fmt.Printf("O valor % .2f é um float\n", i)
	}
}
