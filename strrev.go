package piscine

func StrRev(s string) string {

	bytes := []byte(s)
	var abc int = 0
	var numb byte
	for v := range bytes {
		abc++
		v = v
	}
	for i := 0; i < abc/2; i++ {
		numb = bytes[i]
		bytes[i] = bytes[abc-i-1]
		bytes[abc-i-1] = numb
	}
	return string(bytes)
}
