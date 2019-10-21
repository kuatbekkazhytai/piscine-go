package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb < 20 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}

}
func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
