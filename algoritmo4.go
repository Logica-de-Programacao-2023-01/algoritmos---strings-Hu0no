package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira um texto: ")
	scanner.Scan()
	t1 := scanner.Text()
	fmt.Println("Insira um texto: ")
	scanner.Scan()
	t2 := scanner.Text()
	if t1 == t2 {
		fmt.Println("Os dois textos inseridos são iguais.")
	} else {
		fmt.Print("Os dois textos inseridos são diferentes.")
	}
}
