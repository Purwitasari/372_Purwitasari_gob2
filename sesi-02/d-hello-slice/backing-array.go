package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fmt.Printf("Fruits1  => %#v \n", fruits1)

	var fruits2 = fruits1[2:4]

	fruits2[0] = "rambutan"
	fmt.Printf("Fruits2 => %#v \n", fruits2)

}
