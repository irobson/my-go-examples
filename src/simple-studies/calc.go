package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: n1 n2 op")
		os.Exit(1)
	}
	n1, _ := strconv.ParseInt(os.Args[1], 0, 32)
	n2, _ := strconv.ParseInt(os.Args[2], 0, 32)
	op := os.Args[3]

	if op == "+" {
		fmt.Printf("Soma de %d + %d é: %d\n", n1, n2, n1+n2)
	}
	 if op == "-" {
		fmt.Printf("Subtração de %d - %d é: %d\n", n1, n2, n1-n2)
	}
}
