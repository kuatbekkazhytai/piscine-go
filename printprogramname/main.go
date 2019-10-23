package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args[0]

	for i, j := range arguments {
		z01.PrintRune(j)
		i = i
	}

	z01.PrintRune('\n')
}
