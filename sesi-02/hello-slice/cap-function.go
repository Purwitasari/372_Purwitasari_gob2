package main

import (
	"fmt"
	"strings"
)

func main() {

	var fruits1 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println(fruits1)
	fmt.Println("Fruits1 cap : ", cap(fruits1))
	fmt.Println("Fruits1 len : ", len(fruits1))

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]
	fmt.Println(fruits2)
	fmt.Println("Fruits2 cap : ", cap(fruits2))
	fmt.Println("Fruits2 len : ", len(fruits2))

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]
	fmt.Println(fruits3)
	fmt.Println("Fruits3 cap : ", cap(fruits3))
	fmt.Println("Fruits3 len : ", len(fruits3))

}
