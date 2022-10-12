package main

import (
	"ass-02/ass2/routers"
)

func main() {
	var PORT = ":1212"

	routers.StartServer().Run(PORT)
}
