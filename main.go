package main

import "fmt"

func main() {
	/* SOAL :
		   1. menampilkan nilai i : 21
		   2. menampilkan tipe data dari variabel i
		   3. menampilkan tanda %
		   4. menampilkan nilai boolean j : true
		   5. menampilkan unicode russia : Я (ya)
		   6. menampilkan nilai base 10 : 21
	     7. menampilkan nilai base 8 : 25
		   8. menampilkan nilai base 16 : f
		   9. menampilkan nilai base 16 : F
		   10. menampilkan unicode karakter Я : U+042F
		   11. var k float64 = 123.456; menampilkan float : 123.456000
		   12. menampilkan float scientific : 1.234560E+02
	*/

	// No.1 : menampilkan nilai i : 21
	i := 21
	fmt.Println(i)

	// No.2 : menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// No.3 : menampilkan tanda %
	fmt.Printf("%%\n")

	// No.4 : menampilkan nilai boolean j : true
	j := true
	fmt.Printf("%t\n\n", j)

	// No. ?
	fmt.Printf("%b\n", 21)

	// No.5 : menampilkan unicode russia : Я (ya)
	symbolReverseR := 'Я'
	fmt.Printf("%c\n", symbolReverseR)

	// No.6 : menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", 21)

	// No.7 : menampilkan nilai base 8 : 25
	fmt.Printf("%o\n", 21)

	// No.8 : menampilkan nilai base 16 : f
	fmt.Printf("%x\n", 15)

	// No.9 : menampilkan nilai base 16 : F
	fmt.Printf("%X\n", 15)

	// No.10 : menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U\n", symbolReverseR)

	fmt.Println()

	// No.11 : var k float64 = 123.456; menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%f\n", k)

	// No.12 : menampilkan float scientific : 1.234560E+02
	fmt.Printf("%E\n", k)
}
