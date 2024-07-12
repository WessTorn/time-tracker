package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	DBAdress       string
	DBPort         string
	DBUser         string
	DBPass         string
	DBDatabase     string
	HostURL        string
	ExternalApiURL string
}

var data Config

func InitConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("Failed to load .env file:", err)
	}

	data = Config{
		LogLevel:       os.Getenv("LOG_LEVEL"),
		DBAdress:       os.Getenv("DB_ADDRESS"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASSWORD"),
		DBDatabase:     os.Getenv("DB_DATABASE"),
		HostURL:        os.Getenv("HOST_URL"),
		ExternalApiURL: os.Getenv("EXTERNAL_API_URL"),
	}
}

func LogLevel() string {
	return data.LogLevel
}

func DBAddress() string {
	return data.DBAdress
}

func DBPort() string {
	return data.DBPort
}

func DBUser() string {
	return data.DBUser
}

func DBPass() string {
	return data.DBPass
}

func DBDatabase() string {
	return data.DBDatabase
}

func HostURL() string {
	return data.HostURL
}

func ExternalApiURL() string {
	return data.ExternalApiURL
}
