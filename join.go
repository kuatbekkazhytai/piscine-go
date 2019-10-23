package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i := range strs {

		if i != 0 {
			result = result + sep
		}
		result = result + strs[i]
	}
	return result
}
