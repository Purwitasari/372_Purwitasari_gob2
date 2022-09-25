package main

import "fmt"

func main() {
	greet("Wita", 22)
}

func greet(name string, age int) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}
