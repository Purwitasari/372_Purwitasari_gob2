package main

import "fmt"

func main() {
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	fmt.Printf("Output v interface : %d\n", v)
}
