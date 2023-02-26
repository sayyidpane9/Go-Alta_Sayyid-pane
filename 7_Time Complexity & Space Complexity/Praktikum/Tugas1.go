package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fungsi untuk menghitung hasil dari x^y % m dengan cepat
func fastModExp(x, y, m int) int {
	res := 1
	for y > 0 {
		if y%2 == 1 {
			res = (res * x) % m
		}
		x = (x * x) % m
		y /= 2
	}
	return res
}

// fungsi untuk mengecek apakah sebuah bilangan adalah bilangan prima atau bukan
func primeNumber(number int) bool {
	// bilangan 2 dan 3 adalah bilangan prima
	if number == 2 || number == 3 {
		return true
	}

	// bilangan negatif, 0, dan 1 bukan bilangan prima
	if number <= 1 || number%2 == 0 {
		return false
	}

	// menghitung nilai d dan s pada bilangan number-1
	d, s := number-1, 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	// melakukan uji Miller-Rabin sebanyak 10 kali untuk menghasilkan probabilitas kesalahan 1/4^10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Intn(number-4) + 2
		x := fastModExp(a, d, number)
		if x == 1 || x == number-1 {
			continue
		}
		for j := 0; j < s-1; j++ {
			x = (x * x) % number
			if x == number-1 {
				break
			}
		}
		if x != number-1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
