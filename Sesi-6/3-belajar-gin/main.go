package main

import "3-belajar-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
