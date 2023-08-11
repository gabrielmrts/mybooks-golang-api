package migrations

import (
	"log"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database"
)

func Migrate() {
	db := database.GetDB()

	err := db.AutoMigrate(&models.User{}, &models.Book{})

	if err != nil {
		log.Fatal("Error running migrations")
	}

}
