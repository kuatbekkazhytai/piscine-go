package main

import "fmt"

func StrLen(str string) int {
		change := []rune(str)
		var countlen int = 0
		for i := range change {
			countlen++
			i = i
		}
return countlen
}