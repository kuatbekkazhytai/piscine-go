package piscine

func IsUpper(str string) bool {

	change := []rune(str)
	for i := range change {

		if !(change[i] >= 65 && change[i] <= 90) {
			return false
		}
	}
	return true
}
