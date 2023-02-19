package main

import "fmt"

func main() {
    var sisiAtas, sisiBawah, tinggi float64

    // Input nilai sisi atas, sisi bawah, dan tinggi dari user
    fmt.Println("Masukkan nilai sisi atas trapesium: ")
    fmt.Scanln(&sisiAtas)
    fmt.Println("Masukkan nilai sisi bawah trapesium: ")
    fmt.Scanln(&sisiBawah)
    fmt.Println("Masukkan nilai tinggi trapesium: ")
    fmt.Scanln(&tinggi)

    // Hitung luas trapesium
    luas := 0.5 * (sisiAtas + sisiBawah) * tinggi

    // Tampilkan hasil perhitungan
    fmt.Println("Luas trapesium adalah:", luas)
}
