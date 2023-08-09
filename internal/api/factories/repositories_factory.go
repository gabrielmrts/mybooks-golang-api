package factories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/repositories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database"
)

func GetUsersRepository() *repositories.UserRepository {
	db := database.GetDB()
	return repositories.NewUserRepository(db)
}

func GetBooksRepository() *repositories.BookRepository {
	db := database.GetDB()
	return repositories.NewBookRepository(db)
}
