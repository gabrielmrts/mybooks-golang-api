package database

import (
	"fmt"
	"log"

	"github.com/gabrielmrts/mybooks-golang-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	conf := config.GetConfig()

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s host=%s",
			conf.DATABASE_USER, conf.DATABASE_PASSWORD, conf.DATABASE_NAME, conf.DATABASE_PORT, conf.DATABASE_HOST,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while initializing database")
	}

	log.Println("Database connected!")
	db = database
}

func GetDB() *gorm.DB {
	return db
}
