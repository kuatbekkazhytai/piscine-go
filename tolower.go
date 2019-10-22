package piscine

func ToLower(s string) string {
	change := []rune(s)
	for i, v := range change {
		if change[i] >= 65 && change[i] <= 90 {
			change[i] = change[i] + 32
			v = v
		}
	}
	return string(change)

}
