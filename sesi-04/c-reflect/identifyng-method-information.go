package main

import (
	"fmt"
	"reflect"
)

func (s *student) SetName(Name string) {
	s.Name = Name
}

func main() {
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("Nama: ", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})
	fmt.Println("Nama : ", s1.name)
}
