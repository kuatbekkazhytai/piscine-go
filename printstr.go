package main

import "github.com/01-edu/z01"

func PrintStr(str string) {

stringchange := []rune(str)
			for v := range stringchange{
			z01.Println(stringchange[v])
            }
}