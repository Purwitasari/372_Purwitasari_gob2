package main

import "fmt"

func main() {

	var currentyear = 2021
	// var age sudah tidak perlu karena sudah di inisialisasi dengan :=

	if age := currentyear - 1998; age < 17 {
		fmt.Printf("Kamu belum boleh membuat kartu sim => %d", age)
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

}
