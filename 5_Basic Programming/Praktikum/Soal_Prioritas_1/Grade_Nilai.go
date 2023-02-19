package main

import "fmt"

func main() {
    // Meminta input nilai dari pengguna
    var nilai int
    fmt.Print("Masukkan nilai Anda: ")
    fmt.Scan(&nilai)

    // Memeriksa apakah nilai valid atau tidak
    if nilai < 0 || nilai > 100 {
        fmt.Println("Nilai Invalid")
    } else {
        // Menentukan grade berdasarkan rentang nilai
        var grade string
        switch {
        case nilai >= 80:
            grade = "A"
        case nilai >= 65:
            grade = "B"
        case nilai >= 50:
            grade = "C"
        case nilai >= 35:
            grade = "D"
        default:
            grade = "E"
        }

        // Menampilkan grade
        fmt.Printf("Grade Anda adalah: %s\n", grade)
    }
}
