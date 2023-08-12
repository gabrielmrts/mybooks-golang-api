package migrations

import (
	"log"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/enums"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/helpers"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/google/uuid"
)

func Seed() {
	booksRepository := factories.GetBooksRepository()
	accountsRepository := factories.GetAccountsRepository()

	// Create Accounts
	account := models.Account{
		Email:    "user@email.com",
		Password: helpers.GeneratePasswordHash("123456"),
	}
	user := models.User{
		Name: "user",
		Role: enums.ROLES.USER,
	}
	if err := accountsRepository.Create(&account, &user); err != nil {
		log.Fatal("Failed to seed user: ", err)
	}

	accountAdmin := models.Account{
		Email:    "admin@email.com",
		Password: helpers.GeneratePasswordHash("123456"),
	}
	userAdmin := models.User{
		Name: "admin",
		Role: enums.ROLES.ADMIN,
	}
	if err := accountsRepository.Create(&accountAdmin, &userAdmin); err != nil {
		log.Fatal("Failed to seed user: ", err)
	}

	// Create Books
	books := []models.Book{
		{Title: "book1", ISBN: uuid.NewString(), Price: 20, UserId: 1},
		{Title: "book2", ISBN: uuid.NewString(), Price: 25, UserId: 1},
	}
	for _, book := range books {
		if err := booksRepository.Create(&book); err != nil {
			log.Fatal("Failed to seed book: ", err)
		}
	}

	log.Println("Seed completed")
}
