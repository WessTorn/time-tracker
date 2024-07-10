package main

import (
	"time-tracker/config"
	"time-tracker/database"
	"time-tracker/logger"
	"time-tracker/routers"
)

func main() {
	config.InitConfig()

	logger.InitLogger()

	db := database.ConnectDB()
	defer db.Close()

	database.CreateSchema(db)

	router := routers.InitRouter(db)

	router.Run(config.HostURL())
}
