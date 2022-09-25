package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{}

	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curriculum Developer"

	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee1 : %+v\n", employee1)
	fmt.Printf("Employee2 : %+v\n", employee2)

}
