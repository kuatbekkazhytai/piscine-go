package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args
	for i := range arguments {
		if i != 0 {
			for b, j := range os.Args[i] {
				z01.PrintRune(j)
				b = b
			}
			z01.PrintRune('\n')
		}
	}

}
