package piscine

func IsAlpha(str string) bool {

	change := []rune(str)
	for i := range change {

		if !((change[i] >= '0' && change[i] <= '9') || change[i] >= 97 && change[i] <= 122 || change[i] >= 65 && change[i] <= 90) {
			return false
		}
	}
	return true
}
