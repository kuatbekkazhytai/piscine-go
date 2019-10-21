package piscine

func Sqrt(nb int) int {

	if nb < 0 {
		return 0
	} else {
		result := 0
		for i := 0; i*i <= nb; i++ {
			result = i
		}
		if result*result == nb {
			return result
		} else {
			return 0
		}

	}
}
