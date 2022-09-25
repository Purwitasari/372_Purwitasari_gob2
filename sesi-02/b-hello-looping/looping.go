package main

import "fmt"

func main() {

	for j := 0; j < 3; j++ {
		fmt.Println("Angka", j)
	}

	for j := 1; j <= 10; j++ {
		if j%2 == 1 {
			continue
		}
		if j > 8 {
			break
		}
		fmt.Println("Angka", j)
	}
}
