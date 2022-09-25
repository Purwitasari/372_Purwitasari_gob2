package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
