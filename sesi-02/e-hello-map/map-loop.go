package main

import "fmt"

func main() {

	var person = map[string]string{

		"name":    "Wita",
		"age":     "22",
		"address": "Jakarta",
	}

	for key, value := range person {

		fmt.Println(key, ":", value)
	}

}
