package main

import "fmt"

func countVowels(s string) int {
    count := 0
    for _, c := range s {
        switch c {
            case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
                count++
        }
    }
    return count
}

func main() {
    c1 := "alterra academy"
    c2 := "stringify"
    fmt.Println(countVowels(c1))
    fmt.Println(countVowels(c2))
}
