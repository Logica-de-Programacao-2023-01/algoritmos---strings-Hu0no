package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Print("Insira uma palavra: ")
	fmt.Scan(&x)
	letters := []rune(x)
	LetterCounter := make(map[rune]int)
	for _, letter := range letters {
		LetterCounter[letter]++
	}
	for _, letter := range letters {
		if LetterCounter[letter] == 1 {
			fmt.Print(string(letter))
		}
	}
}
