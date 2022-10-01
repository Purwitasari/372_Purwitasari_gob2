package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	//index/name
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {

	msg := "Heloo World"
	fmt.Fprint(w, msg)

}

/*
 Jalankan dengan perintah go run web-server1.goconst
 Arahkan kepada http://localhost:8080/
 Maka akan keluar tulisan "Hello World"
*/
