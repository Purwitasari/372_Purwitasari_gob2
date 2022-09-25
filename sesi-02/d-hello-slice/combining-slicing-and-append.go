package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fmt.Printf("Fruits1  => %#v \n", fruits1)
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("Fruits1 dengan Rambutan di nomor 4 => %#v \n", fruits1)

}
