package piscine

func IterativePower(nb int, power int) int {

	result := nb
	if power > 0 {

		for i := 1; i <= power-1; i++ {
			result = result * nb
		}
		return result
	} else if power == 0 {
		return 1
	} else {
		return 0
	}

}

