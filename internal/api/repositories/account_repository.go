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

func (ar *AccountRepository) FindByUserID(userId uint) (models.Account, error) {
	var account = models.Account{}

	if err := ar.db.Where(models.Account{UserId: userId}).First(&account).Error; err != nil {
		return models.Account{}, err
	}
	return account, nil
}

func (ar *AccountRepository) SetEmailVerified(account models.Account, email string) error {
	if err := ar.db.Model(&account).Update("EmailVerified", true).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AccountRepository) FindByEmail(email string) (models.Account, error) {
	var account = models.Account{}
	if err := ar.db.Where("email = ?", email).First(&account).Error; err != nil {
		return models.Account{}, err
	}
	return account, nil
}
