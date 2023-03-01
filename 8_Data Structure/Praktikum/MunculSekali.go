package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int)
	for _, digit := range angka {
		num, _ := strconv.Atoi(string(digit))
		counts[num]++
	}
	var result []int
	for num, count := range counts {
		if count == 1 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))			// [4]
	fmt.Println(munculSekali("76523752"))			// [6, 3]
	fmt.Println(munculSekali("12345"))				// [1, 2, 3, 4, 5]
	fmt.Println(munculSekali("1122334455"))			// []
	fmt.Println(munculSekali("0872504"))			// [8 7 2 5 4]
	fmt.Println(" ")
}