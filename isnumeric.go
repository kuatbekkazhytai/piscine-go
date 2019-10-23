package piscine

func IsNumeric(str string) bool {

	change := []rune(str)
	for i := range change {

		if !(change[i] >= '0' && change[i] <= '9') {
			return false
		}
	}
	return true
}
