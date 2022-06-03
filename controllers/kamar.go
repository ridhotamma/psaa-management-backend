package controllers

import (
	"net/http"
	"psaa-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateKamar(c *gin.Context) {
	var input models.KamarCreateInput
	db := c.MustGet("db").(*gorm.DB)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	newKamar := models.Kamar{
		Name:     input.Name,
		Capacity: input.Capacity,
	}

	err = db.Create(&newKamar).Error
	if err != nil {
		c.JSON(201, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newKamar,
	})
}

func UpdateKamar(c *gin.Context) {
	var input models.KamarUpdateInput
	var kamar models.Kamar

	db := c.MustGet("db").(*gorm.DB)

	id, _ := c.Params.Get("id")
	err := db.Where("id = ?", id).First(&kamar).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var anakPanti []models.AnakPanti

	err = db.Find(&anakPanti, input.Members).Error

	if err != nil {
		c.JSON(201, gin.H{
			"error": err.Error(),
		})
	}

	updatedKamar := models.Kamar{
		Name:     input.Name,
		Capacity: input.Capacity,
		Members:  anakPanti,
	}

	err = db.Updates(&kamar).Error
	if err != nil {
		c.JSON(201, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedKamar,
	})
}

func FindKamarByID(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var kamar models.Kamar
	err := db.Where("id = ?", id).First(&kamar).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Data tidak ditemukan!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": kamar,
	})
}

func DeleteKamar(c *gin.Context) {
	id, _ := c.Params.Get("id")
	db := c.MustGet("db").(*gorm.DB)

	var kamar models.Kamar
	err := db.Where("id = ?", id).First(&kamar).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Kamar tidak ditemukan",
		})
		return
	}

	err = db.Delete(&kamar).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": kamar,
	})
}

func ListKamar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var kamar []models.Kamar

	db.Preload(clause.Associations).Find(&kamar)

	c.JSON(http.StatusOK, gin.H{
		"data": kamar,
	})
}
