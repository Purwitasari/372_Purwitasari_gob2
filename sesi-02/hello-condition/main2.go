package main

import "fmt"

func main() {

	var score = 8

	//if jika dari, case => score==10
	if score > 7 {
	}
	switch score {
	case score 10:
		fmt.Println("Perfect")
	default:
		fmt.Println("Nice!")
	}
} else {
	if score ==5 {
		fmt.Println("Not Bad")
	} else if score ==3 {
		fmt.Println("Keep trying")
	} else {
		fmt.Println("You can Do It!")
		if score ==0 {
			fmt.Println("Try Harder!")
		}
	}
}
}
