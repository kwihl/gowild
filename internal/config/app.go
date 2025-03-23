package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gowild.com/pkg/db"
)

type AppEnvironment string

const (
	Local AppEnvironment = "local"
	Test  AppEnvironment = "test"
	Stage AppEnvironment = "stage"
	Prod  AppEnvironment = "prod"
)

type AppConfig struct {
	AppEnvironment
	AppHost  string
	AppPort  string
	DBConfig db.DBConfig
}

func ReadAppConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	appEnvironment := os.Getenv("APP_ENVIRONMENT")
	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return AppConfig{
		AppEnvironment: AppEnvironment(appEnvironment),
		AppHost:        appHost,
		AppPort:        appPort,
		DBConfig: db.DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			DBName:   dbName,
			Username: dbUserName,
			Password: dbPassword,
		},
	}
}
