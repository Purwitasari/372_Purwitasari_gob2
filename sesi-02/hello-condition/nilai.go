package main

import "fmt"

func main() {

	var score = 8

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}
}
