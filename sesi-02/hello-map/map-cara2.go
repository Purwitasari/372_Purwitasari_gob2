package main

import "fmt"

func main() {

	var person = map[string]string{

		"name":    "Wita",
		"age":     "22",
		"address": "Jakarta",
	}

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

}
