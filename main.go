package main

import (
	"fmt"
	"tz_iul/database"
	"tz_iul/routers"
)

func main() {
	db := database.ConnectDB()
	if db == nil {
		fmt.Println("Failed to connect to database")
	}
	defer db.Close()

	database.CreateSchema(db)

	router := routers.InitRouter(db)

	router.Run("localhost:8080")
}
