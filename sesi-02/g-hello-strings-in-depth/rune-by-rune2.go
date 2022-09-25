package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str1 := "makan"
	str2 := "mâkan"

	fmt.Printf("str1 character lenght => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character lenght => %d\n", utf8.RuneCountInString(str2))
}
