package main

import (
	"github.com/WessTorn/time-tracker/config"
	"github.com/WessTorn/time-tracker/database"
	"github.com/WessTorn/time-tracker/logger"
	"github.com/WessTorn/time-tracker/routers"
)

//	@title			Time tracker
//	@version		1.0.0
//	@description	This is an example of a time tracking API..

// @host		localhost:8080
// @BasePath	/
func main() {
	config.InitConfig()
	logger.InitLogger()

	logger.Log.Info("Start")

	logger.Log.Debug("Config:")
	logger.Log.Debug("LOG_LEVEL=" + config.LogLevel())
	logger.Log.Debug("DB_ADDRESS=" + config.DBAddress())
	logger.Log.Debug("DB_PORT=" + config.DBPort())
	logger.Log.Debug("DB_USER=" + config.DBUser())
	logger.Log.Debug("DB_PASSWORD=" + config.DBPass())
	logger.Log.Debug("DB_DATABASE=" + config.DBDatabase())
	logger.Log.Debug("HOST_URL=" + config.HostURL())
	logger.Log.Debug("EXTERNAL_API_URL=" + config.ExternalApiURL())

	db := database.ConnectDB()
	defer db.Close()

	database.CreateSchema(db)

	router := routers.InitRouter(db)

	logger.Log.Infof("(Run) %s", config.HostURL())

	router.Run(config.HostURL())
}
