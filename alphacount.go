package piscine

func AlphaCount(str string) int {

	count := 0
	stringch := []rune(str)
	for index, word := range stringch {
		if word >= 65 && word <= 90 || word >= 97 && word <= 122 {
			count++
			index = index
		}
	}
	return count
}
