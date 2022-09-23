package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fmt.Printf("Fruits1 => %#v \n", fruits1)

	var fruits2 = fruits1[1:4]
	fmt.Printf("Fruits2 1-4 => %#v \n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("Fruits3 0-.. => %#v \n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("Fruits4 ..-3 => %#v \n", fruits4)

	var fruits5 = fruits1[:]
	fmt.Printf("Fruits5 semua => %#v \n", fruits5)

}
