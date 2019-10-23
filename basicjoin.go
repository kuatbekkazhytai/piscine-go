package piscine

func BasicJoin(strs []string) string {

	result := ""
	for i := range strs {
		result = result + strs[i]

	}
	return result
}
