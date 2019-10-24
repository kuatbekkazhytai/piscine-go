package piscine

func Capitalize(s string) string {
	change := []rune(s)
	isFirstLetter := true
	for i := range change {
		if change[i] >= 97 && change[i] <= 122 && isFirstLetter {
			change[i] -= 32
			isFirstLetter = false
		} else if (change[i] >= '0' && change[i] <= '9') || (change[i] >= 65 && change[i] <= 90) && isFirstLetter {
			isFirstLetter = false
		} else if !isFirstLetter && change[i] >= 65 && change[i] <= 90 {
			change[i] += 32
		} else if !isFirstLetter && !(change[i] >= '0' && change[i] <= '9') && !(change[i] >= 97 && change[i] <= 122 || change[i] >= 65 && change[i] <= 90) {
			isFirstLetter = true
		}
	}
	return string(change)
}
