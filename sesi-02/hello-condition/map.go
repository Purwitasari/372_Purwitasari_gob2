package main

import "fmt"

func main() {

	var person map[string]string // Deklarasi

	person = map[string]string{} // Inisialisasi

	person["name"] = "Wita"

	person["age"] = "22"

	person["address"] = "Jalan Sudirman"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

}
