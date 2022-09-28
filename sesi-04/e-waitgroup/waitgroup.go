package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "mango", "durian", "rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}
	wg.Wait()
}

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
