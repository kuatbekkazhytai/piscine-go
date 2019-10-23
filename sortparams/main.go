package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	len := 0
	t := ""
	for i := range os.Args {
		length = i
	}
	for i := 1; i <= len; i++ {
		for j := i + 1; j <= len; j++ {
			if Compare(os.Args[i], os.Args[j]) == 1 {
				t = os.Args[i]
				os.Args[i] = os.Args[j]
				os.Args[j] = t
			}
		}
		PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}
	
}
