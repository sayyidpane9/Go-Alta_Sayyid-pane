package main

import (
	"fmt"
	"strings"
)

func capitalize(str string) string {
	words := strings.Fields(str) 
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, " ") // gabungkan kata-kata kembali menjadi string
}

func main() {
	fmt.Println(capitalize("alterra academy")) 
	fmt.Println(capitalize("ruMaH MAkAn"))     
}
