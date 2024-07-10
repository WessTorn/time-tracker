package main

import (
	"time-tracker/database"
	"time-tracker/logger"
	"time-tracker/routers"
)

func main() {
	logger.InitLogger()

	db := database.ConnectDB()
	defer db.Close()

	database.CreateSchema(db)

	router := routers.InitRouter(db)

	router.Run("localhost:8080")
}
