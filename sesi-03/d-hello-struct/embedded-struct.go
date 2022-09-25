package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {

	var employee1 = Employee{}

	employee1.person.name = "Airell"
	employee1.person.age = 23
	employee1.division = "Curriculum Developer"

	fmt.Printf("%+v", employee1)
}
