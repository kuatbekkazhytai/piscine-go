package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i := range strs {
		res = res + (strs[i] + sep)
	}
	return res
}
