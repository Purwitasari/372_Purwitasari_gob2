package main

import "fmt"

func main() {

	var person = map[string]string{

		"name":    "Wita",
		"age":     "22",
		"address": "Jakarta",
	}

	fmt.Println("Before deleting:", person)
	delete(person, "age")
	fmt.Println("After deleting:", person)

}
