package controllers

import (
	"gin-bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBookInput - Input to create book
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput - Input to update book
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks - Method to return list of books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Order("created_at desc").Find(&books)

	c.JSON(http.StatusOK, books)
}

// FindBook - Method to return a book
func FindBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook - Method to create a book
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, book)
}

// UpdateBook - Method to update a book
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, book)
}

// DeleteBook - Method to delete a book
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, book)
}
