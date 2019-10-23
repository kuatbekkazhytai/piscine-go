package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args
	length := 0
	for i := range arguments {
		length = i
	}
	for i := length; i > 0; i-- {
		for b, j := range os.Args[i] {
			z01.PrintRune(j)
			b = b
		}
		z01.PrintRune('\n')
	}

}
