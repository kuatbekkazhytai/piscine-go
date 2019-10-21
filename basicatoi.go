package piscine


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
