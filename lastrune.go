package piscine

func LastRune(s string) rune {

	change := []rune(s)
	var n int = 0

	for i := range change {
		n++
		i = i
	}
	return NRune(s, n)

}
