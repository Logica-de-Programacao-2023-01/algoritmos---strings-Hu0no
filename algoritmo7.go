package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insira algum texto: ")
	scanner.Scan()
	t := scanner.Text()
	x := strings.Split(t, " ")
	i := 0
	for i < len(x) {
		y, err := strconv.ParseFloat(x[i], 64)
		if err == nil && y == y {
			fmt.Print("Há números na string")
			break
		}
		i++
	}
}
