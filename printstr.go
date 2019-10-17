package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!"
	PrintStr(str)
}


func PrintStr(str string) {

stringchange := []rune(str)
for v := range stringchange{
z01.Println(stringchange[v])

}
}
