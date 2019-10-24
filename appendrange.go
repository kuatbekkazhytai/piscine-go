package piscine

func AppendRange(min, max int) []int {

	var ans []int

	for i := min - 1; i < max-1; i++ {
		ans = append(ans, i+1)

	}

	return ans
}
