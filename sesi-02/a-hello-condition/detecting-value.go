package main

import "fmt"

func main() {

	var person = map[string]string{

		"name":    "Wita",
		"age":     "22",
		"address": "Jakarta",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println(("Value does'nt exist"))
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println(("Value has been deleted"))
	}

}
