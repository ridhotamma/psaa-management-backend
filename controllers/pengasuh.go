package controllers

import (
	"net/http"
	"psaa-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePengasuh(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.PengasuhCreateInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

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

	newPengasuh := models.Pengasuh{
		Kamar: kamar,
		User:  user,
	}

	err = db.Create(&newPengasuh).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newPengasuh,
	})
}
