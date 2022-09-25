package main

import "fmt"

func main() {

	var score = 8

	//if jika dari, case => score==10
	if score > 7 {
	}
	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad")
		fallthrough
	case score < 5:
		fmt.Println(("It is ok, but please harder"))
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}
}
