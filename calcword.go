package main

import "fmt"

func main() {
	input := "selamat malam"

	charCount := make(map[string]int)

	for _, s := range input {
		fmt.Printf("%c\n", s)
		charCount[string(s)] += 1
	}

	fmt.Println(charCount)
}
