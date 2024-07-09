package main

import "tz_iul/routers"

func main() {
	router := routers.InitRouter()

	router.Run("localhost:8080")
}
