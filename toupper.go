package piscine

func ToUpper(s string) string {
	change := []rune(s)
	for i, v := range change {
		if change[i] >= 97 && change[i] <= 122 {
			change[i] = change[i] - 32
			v = v
		}
	}
	return string(change)
}
