package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}
	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name : ", employee1.name)
	fmt.Println("Employee2 name : ", employee2.name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.name = "Ananda"

	fmt.Println("Employee1 name : ", employee1.name)
	fmt.Println("EMployee2 name : ", employee2.name)

}
