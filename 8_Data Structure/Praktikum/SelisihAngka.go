package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	var sumLeft, sumRight int32

	// Hitung jumlah diagonal dari kiri ke kanan
	for i := 0; i < len(arr); i++ {
		sumLeft += arr[i][i]
	}

	// Hitung jumlah diagonal dari kanan ke kiri
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		sumRight += arr[i][j]
	}

	// Hitung selisih mutlak antara kedua diagonal
	diff := int32(math.Abs(float64(sumLeft - sumRight)))

	return diff
}

func main() {
	// Contoh input matriks
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	// Hitung selisih mutlak antara diagonal
	diff := diagonalDifference(arr)

	fmt.Println("Selisihnya adalah : ", diff) // Output: 2
	fmt.Println(" ")
}
