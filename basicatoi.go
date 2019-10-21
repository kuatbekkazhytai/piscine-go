package piscine

import (
        "fmt"
)

func main() {
        s := "12345"
        s2 := "0000000012345"
        s3 := "000000"

        n := BasicAtoi(s)
        n2 := BasicAtoi(s2)
        n3 := BasicAtoi(s3)

        fmt.Println(n)
        fmt.Println(n2)
        fmt.Println(n3)
}
func BasicAtoi(s string) int {
        change := []rune(s)
        var mean int = 0
        for _, r := range change {
                mean *= 10
                if r == '0' {
                        mean += 0
                } else if r == '1' {
                        mean += 1
                } else if r == '2' {
                        mean += 2
                } else if r == '3' {
                        mean += 3
                } else if r == '4' {
                        mean += 4
                } else if r == '5' {
                        mean += 5
                } else if r == '6' {
                        mean += 6
                } else if r == '7' {
                        mean += 7
                } else if r == '8' {
                        mean += 8
                } else if r == '9' {
                        mean += 9
                }
        }
        return mean
}
