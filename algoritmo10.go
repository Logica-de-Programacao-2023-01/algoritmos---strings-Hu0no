package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var t2 string

func alphaOnly(t2) bool {
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func main() {
	var i int = 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira uma palavra: ")
	scanner.Scan()
	t1 := scanner.Text()
	fmt.Println("Insira outra palavra: ")
	scanner.Scan()
	t2 = scanner.Text()
	for {
		if len(t1) != len(t2) {
			fmt.Print("As duas palavras não são anagramas.")
			break
		} else {
			for i < len(t1) {
				if IsLetter(t2) == false {
					print("As duas palavras não são anagramas.")
					break
				}
				i++
			}
		}
		print("As duas palavras são anagramas.")
	}

}
