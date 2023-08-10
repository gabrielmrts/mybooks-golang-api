package repositories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) List() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Preload("Books").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	err := ur.db.Create(user).Error
	return err
}

func (ur *UserRepository) FindByEmail(email string) (models.User, error) {
	var user = models.User{}
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) FindByEmailAndPassword(email string, password string) (models.User, error) {
	var user = models.User{}
	if err := ur.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
