package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateBookRequestBody struct {
	Title string  `json:"title" binding:"required,max=255"`
	ISBN  string  `json:"isbn" binding:"required,numeric,min=10,max=10"`
	Price float64 `json:"price" binding:"required"`
}

func ListBooks(c *gin.Context) {
	booksRepository := factories.GetBooksRepository()
	books, err := booksRepository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var requestBody CreateBookRequestBody
	booksRepository := factories.GetBooksRepository()

	authorIdAny, _ := c.Get("user_id")
	authorId, _ := authorIdAny.(uint)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	var book = models.Book{Title: requestBody.Title, ISBN: requestBody.ISBN, Price: requestBody.Price, UserId: int(authorId)}
	if err := booksRepository.Create(&book); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Duplicated ISBN"})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error while creating a book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created with success"})
}
