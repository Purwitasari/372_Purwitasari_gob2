package main

import "fmt"

func main() {

	//multiple declaration
	var name, age, address = "Wita", 23, "Jakarta"
	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address)
	fmt.Println(first, second, third)

	var firstVariable string

	/*
		Tidak ada variable yang boleh menganggur ketika sudah dibuat
		Variable underscore digunakan untuk menghindari error
	*/

	_, _, _, _ = firstVariable, name, age, address

	my_first, my_second := 1, "2"
	// Perbedaan penulisan ""
	fmt.Printf("Tipe data variabel first adalah %T \n", my_first)
	fmt.Printf("Tipe data variabel second adalah %T \n", my_second)

	//Menggunakan verb
	fmt.Printf("Hallo nama ku %s, umurku adalah %d, dan aku tinggal di %s", name, age, address)
}
