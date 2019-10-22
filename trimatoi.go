package piscine

func TrimAtoi(s string) int {
	var negative bool = false
	var empty bool = true
	var res int = 0
	for _, v := range s {
		if empty && !negative && v == '-' {
			negative = true
		} else if DigitRune(v) {
			res *= 10
			res += int(v - 48)
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if negative {
			return -res
		} else {
			return res
		}
	}
}