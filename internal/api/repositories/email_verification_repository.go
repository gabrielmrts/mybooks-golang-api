package repositories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"gorm.io/gorm"
)

type EmailVerificationRepository struct {
	db *gorm.DB
}

func NewEmailVerificationRepository(db *gorm.DB) *EmailVerificationRepository {
	return &EmailVerificationRepository{db: db}
}

func (evr *EmailVerificationRepository) Create(ev *models.EmailVerification) error {
	err := evr.db.Create(ev).Error
	return err
}

func (evr *EmailVerificationRepository) GetByCode(code string) (models.EmailVerification, error) {
	var emailVerification = models.EmailVerification{}

	if err := evr.db.Where(models.EmailVerification{Code: code}).First(&emailVerification).Error; err != nil {
		return models.EmailVerification{}, err
	}
	return emailVerification, nil
}
