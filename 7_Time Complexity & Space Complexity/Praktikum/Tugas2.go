package main

import "fmt"

func pow(x, n int) int {
    if n == 0 {
        return 1
    }
    if n%2 == 0 {
        y := pow(x, n/2)
        return y * y
    }
    y := pow(x, (n-1)/2)
    return x * y * y
}

func main() {
    fmt.Println(pow(2, 3))  // output: 8
    fmt.Println(pow(7, 2))  // output: 49
    fmt.Println(pow(5, 3))  // output: 125
    fmt.Println(pow(10, 2)) // output: 100
    fmt.Println(pow(2, 5))  // output: 32
    fmt.Println(pow(7, 3))  // output: 343
}
