package controllers

import (
	"apiGo/database"
	"apiGo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id) //converte para int
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book " + err.Error(),
		})
		return
	}

	c.JSON(201, book)

}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books " + err.Error(),
		})
		return
	}

	c.JSON(200, books)

}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book " + err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id) //converte para int
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleted for success",
	})
}