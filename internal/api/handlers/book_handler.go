package handlers

import (
	"net/http"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/factories"
	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	booksRepository := factories.GetBooksRepository()
	books, err := booksRepository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, books)
}
