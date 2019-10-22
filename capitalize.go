package piscine

func Capitalize(s string) string {
	change := []rune(s)
	FirstLetter := true
	for i := range change {
		if change[i] >= 'a' && change[i] <= 'z' && FirstLetter {
			change[i] -= 32
			FirstLetter = false
		} else if (IsRuneDigit(change[i]) || (change[i] >= 'A' && change[i] <= 'Z')) && FirstLetter {
			FirstLetter = false
		} else if !FirstLetter && change[i] >= 'A' && change[i] <= 'Z' {
			change[i] += 32
		} else if !FirstLetter && !IsRuneDigit(change[i]) && !IsRuneAlpha(change[i]) {
			FirstLetter = true
		}
	}
	return string(change)
}
