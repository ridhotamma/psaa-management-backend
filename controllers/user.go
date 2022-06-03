package controllers

import (
	"net/http"
	"psaa-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var userInput models.UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser := models.User{
		FirstName:   userInput.FirstName,
		LastName:    userInput.LastName,
		Email:       userInput.Email,
		Username:    userInput.Username,
		Password:    userInput.Password,
		PhoneNumber: userInput.PhoneNumber,
	}

	db := c.MustGet("db").(*gorm.DB)
	err = db.Create(&newUser).Error

	if err != nil {
		c.JSON(201, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newUser,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userInput models.UserInput
	err = c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser := models.User{
		FirstName:   userInput.FirstName,
		LastName:    userInput.LastName,
		Email:       userInput.Email,
		Password:    userInput.Password,
		Username:    userInput.Username,
		PhoneNumber: userInput.PhoneNumber,
	}

	db.Model(&user).Updates(updatedUser)
}

func DeleteUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UserList(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func FindUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
