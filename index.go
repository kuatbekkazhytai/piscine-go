package piscine

func Index(s string, toFind string) int {
	len1 := StrLen(s)
	len2 := StrLen(toFind)
	found := true

	if len2 < 1 {
		return 0
	} else if len1 < 1 {
		return -1
	} else {

		rune1 := []rune(s)
		rune2 := []rune(toFind)
		for i := 0; i < len1; i++ {
			found = true
			for j := 0; j < len2; j++ {
				if rune1[i+j] != rune2[j] {
					found = false
					j = len2
				}
			}
			if found {
				return i

			}
		}

		return -1
	}

}
