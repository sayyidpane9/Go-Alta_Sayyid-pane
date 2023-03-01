package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// Buat map untuk menyimpan nilai dan indeks
    numMap := make(map[int]int)

    // Loop melalui array
    for i, num := range arr {
        // Cari selisih antara target dan nilai saat ini
        complement := target - num

        // Cek apakah selisih sudah ada di map
        if _, ok := numMap[complement]; ok {
            // Jika ditemukan, return array indeks pasangan
            return []int{numMap[complement], i}
        }

        // Tambahkan nilai dan indeks ke map
        numMap[num] = i
    }

    // Jika tidak ditemukan, return array kosong
    return []int{}
}

func main() {
    fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))    	// [1 3]
    fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))     	// [0 2]
    fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   		// [2 3]
    fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) 		// [1 2]
    fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) 			// [0 1]
	fmt.Println(" ")
}
