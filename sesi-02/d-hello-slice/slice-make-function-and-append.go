package main

import "fmt"

func main() {
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	/*
		Bisa juga dituliskan dengan

		fruits = append(fruits, "apple", "banana", "mango")

	*/

	fmt.Printf("%#v", fruits)

}
