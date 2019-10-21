package piscine

func Swap(a *int, b *int) {
		tempa := *a
        tempb := *b
        *a = tempa
		*b = tempb
}
