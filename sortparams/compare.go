package piscine

func Compare(a, b string) int {
	len1 := StrLen(a)
	len2 := StrLen(b)
	n := len1
	rune1 := []rune(a)
	rune2 := []rune(b)
	if len1 == 0 && len2 == 0 {
		return 0
	} else if len1 == 0 {
		return -1
	} else if len2 == 0 {
		return 1
	} else {
		if len1 > len2 {
			n = len2
		}
		for i := 0; i < n; i++ {
			if rune1[i] > rune2[i] {
				return 1
			} else if rune1[i] < rune2[i] {
				return -1
			}
		}
		if len1 == len2 {
			return 0
		} else if len1 > len2 {
			return 1
		}
		return -1
	}
}
