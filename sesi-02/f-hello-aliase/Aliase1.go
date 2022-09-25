package main

import "fmt"

func main() {
	// deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari tipe data uint8

	b = a // no error, karena byte memiliki tipe data yang sama dengan uint8
	_ = b
	fmt.Printf("%d, %d", a, b)
}
