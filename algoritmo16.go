package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	var y string
	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&x)
	fmt.Println("Digite outra palavra: ")
	fmt.Scan(&y)
	x = strings.ToLower(x)
	y = strings.ToLower(y)
	z := strings.Contains(x, y)
	if z == true {
		fmt.Println("A segunda palavra está contida na primeira")
	} else {
		fmt.Println("A segunda palavra não está contida na primeira")
	}
}
