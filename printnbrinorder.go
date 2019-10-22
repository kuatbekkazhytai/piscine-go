package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	var x []int
	if n == 0 {
		x = append(x, 0)
	}

	for i := 0; n > 0; i++ {
		x = append(x, n%10)
		n = n / 10

	}
	SortIntegerTable(x)
	for i := range x {
		z01.PrintRune(rune('0' + x[i]))
	}
}
