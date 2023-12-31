package factories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/repositories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database"
)

func GetUsersRepository() *repositories.UserRepository {
	db := database.GetDB()
	return repositories.NewUserRepository(db)
}

func GetEmailVerificationRepository() *repositories.EmailVerificationRepository {
	db := database.GetDB()
	return repositories.NewEmailVerificationRepository(db)
}

func GetBooksRepository() *repositories.BookRepository {
	db := database.GetDB()
	return repositories.NewBookRepository(db)
}

func GetAccountsRepository() *repositories.AccountRepository {
	db := database.GetDB()
	return repositories.NewAccountRepository(db)
}
