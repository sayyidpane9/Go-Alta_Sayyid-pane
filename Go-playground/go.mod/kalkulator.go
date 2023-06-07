package main

import "fmt"

type calculator struct {
}

func (c calculator) add(a, b float64) float64 {
    return a + b
}

func (c calculator) subtract(a, b float64) float64 {
    return a - b
}

func (c calculator) multiply(a, b float64) float64 {
    return a * b
}

func (c calculator) divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("Pembagi 0")
    }
    return a / b, nil
}

func main() {
    c := calculator{}

    fmt.Println("Penambahan:", c.add(10, 5))
    fmt.Println("Pengurangan:", c.subtract(10, 5))
    fmt.Println("Perkalian:", c.multiply(10, 5))

    result, err := c.divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", result)
    }

    result, err = c.divide(10, 5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", result)
    }
	fmt.Println(" ")
}
