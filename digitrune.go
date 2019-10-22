package piscine

func DigitRune(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}