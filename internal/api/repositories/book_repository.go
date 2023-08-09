package repositories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (ur *BookRepository) List() ([]models.Book, error) {
	var books []models.Book
	if err := ur.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (ur *BookRepository) Create(book *models.Book) error {
	err := ur.db.Create(book).Error
	return err
}
