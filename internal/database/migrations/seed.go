package migrations

import (
	"log"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/google/uuid"
)

func Seed() {
	usersRepository := factories.GetUsersRepository()
	booksRepositoy := factories.GetBooksRepository()

	users := []models.User{
		{Name: "user1", Email: "user1@example.com", Password: "user1password"},
	}
	books := []models.Book{
		{Title: "book1", ISBN: uuid.NewString(), Price: 20, UserId: 1},
		{Title: "book2", ISBN: uuid.NewString(), Price: 25, UserId: 1},
	}

	for _, user := range users {
		if err := usersRepository.Create(&user); err != nil {
			log.Fatal("Failed to seed user: ", err)
		}
	}
	for _, book := range books {
		if err := booksRepositoy.Create(&book); err != nil {
			log.Fatal("Failed to seed book: ", err)
		}
	}

	log.Println("Seed completed")
}
