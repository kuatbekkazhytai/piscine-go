package piscine

func IsAlpha(str string) bool {

	change := []rune(str)
	for i := range change {

		if !(change[i] >= 97 && change[i] <= 122) {
			return false
		}
	}
	return true
}
