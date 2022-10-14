package main

import (
	"rest/database"
	"rest/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}
