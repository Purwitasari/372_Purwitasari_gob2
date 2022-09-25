package main

import (
	"fmt"
	"strings"
)

func main() {

	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")

}

func profile(name string, favFoods ...string) {

	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there!!! I'm ", name)
	fmt.Println("I really love to eat ", mergeFavFoods)

}
