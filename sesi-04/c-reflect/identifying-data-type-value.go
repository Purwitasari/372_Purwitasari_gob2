package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Printf("reflect value :  %+v\n", number)
	fmt.Println("tipe variable : ", reflectValue.Type())
	fmt.Printf("nilai variable :  %+v\n", reflectValue.Interface())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable : ", reflectValue.Int())
	}
}
