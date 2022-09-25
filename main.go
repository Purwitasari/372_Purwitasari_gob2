package main

import "belajar/sesi03-e-hello-exported-unexported"

func main() {
	helpers.Greet()
	var person = helpers.Person{}

	person.Invokegreet()
}
