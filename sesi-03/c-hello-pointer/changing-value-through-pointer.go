package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memory address) : ", &firstPerson)

	fmt.Println("secondNumber (value) : ", *secondPerson)
	fmt.Println("secondNumber (memory address) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memory address) : ", &firstPerson)

	fmt.Println("secondNumber (value) : ", *secondPerson)
	fmt.Println("secondNumber (memory address) : ", secondPerson)

}
