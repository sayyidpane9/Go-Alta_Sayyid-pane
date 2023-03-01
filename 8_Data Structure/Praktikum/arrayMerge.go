package main

import "fmt"

// fungsi untuk menggabungkan dua array dan menghindari duplikasi nama
func ArrayMerge(arrayA, arrayB []string) []string {
    // map untuk menyimpan nama yang sudah ada
    existingNames := make(map[string]bool)

    // slice untuk menyimpan hasil gabungan
    result := []string{}

    // tambahkan nama dari array pertama ke hasil gabungan
    for _, name := range arrayA {
        if !existingNames[name] {
            result = append(result, name)
            existingNames[name] = true
        }

	}

    // tambahkan nama dari array kedua ke hasil gabungan
    for _, name := range arrayB {
        if !existingNames[name] {
            result = append(result, name)
            existingNames[name] = true
        }
    }

    // kembalikan hasil gabungan
    return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ("king", "devil jin", "akuma", "eddie", "steve", "geese")

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ("sergei", "jin", "steve", "bryan")

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	// ("alisa", "yoshimitsu", "devil jin", "law")

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin",  "sergei"}))
	// ("devil jin", "sergei")

	fmt.Println(ArrayMerge([]string {"hwoarang"}, []string{}))
	// ("hwoarang")

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
    fmt.Println(" ")
}
