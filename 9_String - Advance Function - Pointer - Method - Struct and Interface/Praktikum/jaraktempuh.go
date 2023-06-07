package main

import "fmt"

type Car struct {
    Type string
    FuelIn float64 // dalam liter
}

func (c *Car) EstimateDistance() float64 {
    var fuelEfficiency float64
    if c.Type == "Fortuner" {
        fuelEfficiency = 14 // km/L
    } else if c.Type == "Pajero" {
        fuelEfficiency = 10 // km/L
    } else {
        fuelEfficiency = 12 // km/L
    }
    return fuelEfficiency * c.FuelIn * 1.5 // 1.5 km/mill
}

func main() {
    fortuner := Car{"Fortuner", 30} // fortuner dengan 30 liter bahan bakar
    pajero := Car{"Pajero", 40} // pajero dengan 40 liter bahan bakar
    sedan := Car{"Sedan Biasa", 25} // sedan biasa dengan 25 liter bahan bakar

    fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh Fortuner\t: %.2f km\n", fortuner.EstimateDistance())
    fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh Pajero\t\t: %.2f km\n", pajero.EstimateDistance())
    fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh Sedan Biasa\t: %.2f km\n", sedan.EstimateDistance())
	fmt.Println(" ")
}
