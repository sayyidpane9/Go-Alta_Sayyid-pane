package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan jumlah : ")
    fmt.Scanln(&n)

    for i := 0; i < n; i++ {
        for j := 0; j < n-i; j++ {
            fmt.Print(" ")
        }
        for k := 0; k <= i; k++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}
