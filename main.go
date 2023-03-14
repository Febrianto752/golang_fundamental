package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		if i < 5 {
			fmt.Println("Nilai i =", i)

		} else {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					const RussianUnicodes = "САШАРВО"
					for index, value := range RussianUnicodes {
						fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index)
					}
				} else {
					fmt.Println("Nilai j =", j)

				}
			}
		}
	}
}
