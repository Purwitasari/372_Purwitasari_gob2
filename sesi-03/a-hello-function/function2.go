package main

import "fmt"

func main() {
	greet("Wita", "Jakarta")
}

func greet(name, address string) {
	fmt.Println("Hello there! My name is ", name)
	fmt.Println("I live in ", address)
}
