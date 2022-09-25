package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firtsNumber (value) : ", firstNumber)
	fmt.Println("firtsNumber (memory address) : ", &firstNumber)

	fmt.Println("secondNumber (value) : ", *secondNumber)
	fmt.Println("secondNumber (memory address) : ", secondNumber)

}
