package main

import "fmt"

func main() {

	str := "mâkan"

	for i, s := range str {
		fmt.Printf("index=> %d, string => %s \n", i, string(s))
	}
}
