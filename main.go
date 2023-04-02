package main

import "simple_rest_api_book/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}