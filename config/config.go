package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	ENVIRONMENT       string
	SMTP_HOST         string
	SMTP_PORT         int
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
	environment := os.Getenv("ENVIRONMENT")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	smtpPortInt, _ := strconv.Atoi(smtpPort)

	config := &Config{
		DATABASE_HOST:     databaseHost,
		DATABASE_PORT:     databasePort,
		DATABASE_NAME:     databaseName,
		DATABASE_USER:     databaseUser,
		DATABASE_PASSWORD: databasePassword,
		ENVIRONMENT:       environment,
		SMTP_HOST:         smtpHost,
		SMTP_PORT:         smtpPortInt,
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
