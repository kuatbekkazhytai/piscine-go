package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {

	change := []rune(s)
	var count int = 0

	for i := range change {
		count ++
		i = i
	}
	if n <= count && count !=0 && n>=0 {
	
		return change[n-1]

	} else {
		return 0
	}
}
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}