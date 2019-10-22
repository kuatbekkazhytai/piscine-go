package piscine

func Capitalize(s string) string {
	change := []rune(s)
	FirstLetter := true
	for i := range runes {
		if change[i] >= 'a' && change[i] <= 'z' && FirstLetter {
			change[i] -= 32
			FirstLetter = false
		} else if (IsRuneDigit(change[i]) || (change[i] >= 'A' && change[i] <= 'Z')) && FirstLetter {
			FirstLetter = false
		} else if !isFirstLetter && change[i] >= 'A' && change[i] <= 'Z' {
			change[i] += 32
		} else if !isFirstLetter && !IsRuneDigit(change[i]) && !IsRuneAlpha(change[i]) {
			FirstLetter = true
		}
	}
	return string(change)
}
