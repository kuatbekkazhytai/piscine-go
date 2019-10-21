package piscine

import "fmt"

func main() {
	str := "Hello World!"
	nb :=StrLen(str)
	fmt.Println(nb)
}

func StrLen(str string) int {

change := []rune(str)
var countlen int = 0
for i := range change {
	countlen++
	i = i
}
return countlen
}
