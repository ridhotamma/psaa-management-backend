package controllers

import (
	"net/http"
	"psaa-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateAnakPanti(c *gin.Context) {
	var input models.AnakPantiCreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	err = db.Where("id = ?", input.UserID).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var kamar models.Kamar
	err = db.Where("id = ?", input.KamarID).First(&kamar).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser := models.AnakPanti{
		User:  user,
		Kamar: kamar,
	}

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

func UpdateAnakPanti(c *gin.Context) {
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

func DeleteAnakPanti(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var anakPanti models.AnakPanti
	err := db.Where("id = ?", id).First(&anakPanti).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Delete(&anakPanti)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func AnakPantiList(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var anakPanti []models.AnakPanti
	db.Preload(clause.Associations).Find(&anakPanti)

	c.JSON(http.StatusOK, gin.H{
		"data": anakPanti,
	})
}

func FindAnakPantiById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var anakPanti models.AnakPanti

	err := db.Preload(clause.Associations).Where("id = ?", id).First(&anakPanti).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": anakPanti,
	})
}
