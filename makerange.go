package piscine

func MakeRange(min, max int) []int {

	answer := []int(nil)
	val := max - min

	if max > min {
		answer = make([]int, val)
		for i := 0; min < max; i++ {
			answer[i] = min
			min++
		}
	}
	return answer
}
