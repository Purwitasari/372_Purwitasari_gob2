package main

import (
	"ass-03/ass3/routers"
)

func main() {
	var PORT = ":1212"

	routers.StartServer().Run(PORT)
}
