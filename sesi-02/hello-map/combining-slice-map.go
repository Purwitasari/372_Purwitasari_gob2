package main

import "fmt"

func main() {

	var people = []map[string]string{
		{"name": "Wita", "age": "22"},
		{"name": "Fauzan", "age": "23"},
		{"name": "Aci", "age": "23"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age %s\n", i, person["name"], person["age"])
	}

}
