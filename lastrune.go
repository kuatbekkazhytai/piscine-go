package piscine

func NRune(s string, n int) rune {

	change := []rune(s)
	var count int = 0

	for i := range change {
		count++
		i = i
	}
	if n <= count && count != 0 && n >= 0 {

		return change[n-1]

	} else {
		return 0
	}
}
func LastRune(s string) rune {

	change := []rune(s)
	var n int = 0

	for i := range change {
		n++
		i = i
	}
	return NRune(s, n)

}
