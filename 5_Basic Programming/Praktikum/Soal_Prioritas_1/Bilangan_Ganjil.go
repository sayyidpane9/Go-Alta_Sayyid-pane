package main

import "fmt"

func main() {
    // Meminta input bilangan dari pengguna
    var bilangan int
    fmt.Print("Masukkan sebuah bilangan: ")
    fmt.Scan(&bilangan)

    // Menggunakan operator modulo untuk menentukan apakah bilangan itu genap atau ganjil
    if bilangan%2 == 0 {
        fmt.Printf("%d adalah bilangan genap\n", bilangan)
    } else {
        fmt.Printf("%d adalah bilangan ganjil\n", bilangan)
    }
}
