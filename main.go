package main

import (
	"fmt"
	"strconv"
)

func main() {
	unicodesValues := [7]string{"\u0421", "\u0410", "\u0428", "\u0410", "\u0420", "\u0412", "\u041E"}
	unicodesString := [7]string{"U+0421", "U+0410", "U+0428", "U+0410", "U+0420", "U+0412", "U+041E"}

	for i := 0; i <= 5; i++ {
		if i < 5 {
			fmt.Println("Nilai i =", i)

		} else {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					var position int64 = 0
					for index, unicode := range unicodesValues {
						fmt.Println("character " + unicodesString[index] + " '" + unicode + "' starts at byte position " + strconv.FormatInt(position, 10))
						position += 2
					}
				} else {
					fmt.Println("Nilai j =", j)

				}
			}
		}
	}
}
