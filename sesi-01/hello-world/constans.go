package main

import "fmt"

func main() {

	//Constant merupakan jenis variable yang nilainya tidak dapat diubah
	const full_name string = "Purwitasari Muliawan"
	fmt.Printf("Hello %s \n", full_name)

	var value = 2 + 2*3
	fmt.Println("Hasil 2 + 2*3 => ", value)

	var my_value = (2 + 2) * 3
	fmt.Println("Hasil (2 + 2) * 3 => ", my_value)

}
