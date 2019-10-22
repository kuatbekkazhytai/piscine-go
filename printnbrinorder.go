package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}

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
func SortIntegerTable(table []int) {
	var length int = 0
	for _, v := range table {
		v = v
		length++
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if table[i] > table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
func Swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}
