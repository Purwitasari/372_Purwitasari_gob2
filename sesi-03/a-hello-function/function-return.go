package main

import (
	"fmt"
	"strings"
)

func main() {

	var names = []string{"Purwitasari", "Muliawan"}
	var printMsg = greet("Heii", names)

	fmt.Println(printMsg)

}

func greet(msg string, names []string) string {

	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result

}
