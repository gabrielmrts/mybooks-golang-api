package repositories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Create(account *models.Account, user *models.User) error {
	// initialize transaction
	tx := ar.db.Begin()

	// try to create a user
	err := ar.db.Create(user).Error

	if err != nil {
		if err != nil {
			if e := tx.Rollback().Error; e != nil {
				return e
			}
		}
	}

	// assign a user to a new account
	account.UserId = user.ID

	// then create the account
	err = ar.db.Create(account).Error

	// if fail then rollback
	if err != nil {
		if e := tx.Rollback().Error; e != nil {
			return e
		}
	}

	// if all ok then commit
	if err = tx.Commit().Error; err != nil {
		return err
	}

	return err
}
