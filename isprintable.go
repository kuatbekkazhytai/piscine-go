package piscine

func IsPrintable(str string) bool {

	change := []rune(str)

	for i := range change {

		if change[i] < 32 {
			return false
		}

	}
	return true

}
