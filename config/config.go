package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
}

func LoadConfig() *Config {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	config := &Config{
		DATABASE_HOST:     databaseHost,
		DATABASE_PORT:     databasePort,
		DATABASE_NAME:     databaseName,
		DATABASE_USER:     databaseUser,
		DATABASE_PASSWORD: databasePassword,
	}

	return config
}

var globalConfig *Config

func Init() {
	globalConfig = LoadConfig()
}

func GetConfig() *Config {
	return globalConfig
}
