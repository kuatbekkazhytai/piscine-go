package piscine

func ConcatParams(args []string) string {

	res := ""

	for i, v := range args {
		if i != 0 {
			res = res + "\n"
		}

		res = res + v

	}
	return res
}
