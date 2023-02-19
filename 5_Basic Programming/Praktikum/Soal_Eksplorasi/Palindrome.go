package main

import (
	"fmt"
	"strings"
)


func main() {	
    var input string
    fmt.Print("Apakah Palindrome?", "\nMasukkan kata : ")
    fmt.Scanln(&input)

    kata := strings.Fields(strings.ToLower(input)) // Memisahkan kata-kata dan mengubah menjadi huruf kecil

    for _, k := range kata {
        palindrome := true
        for i := 0; i < len(k)/2; i++ {
            if k[i] != k[len(k)-i-1] {
                palindrome = false
                break
            }
        }

        if palindrome {
            fmt.Println("Captured : " ,k, "\npalindrome.\n")
        } else {
            fmt.Println("Captured : " ,k, "\nbukan palindrome.\n")
        }
    }
}
