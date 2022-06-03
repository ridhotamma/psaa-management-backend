package controllers

import (
	"log"
	"net/http"
	"psaa-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductInput struct {
	Name       string `json:"name" binding:"required"`
	Categories []int  `json:"categories" binding:"required"`
}

type CategoryInput struct {
	Name string
}

func ListProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product []models.Product
	err := db.Preload(clause.Associations).Find(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func ListCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category []models.Category
	err := db.Find(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CategoryInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category := models.Category{
		Name: input.Name,
	}

	err = db.Create(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var categories []models.Category

	db.Find(&categories, input.Categories)

	log.Printf("category data: %v input data: %v", categories, input.Categories)

	newProduct := models.Product{
		Name:       input.Name,
		Categories: categories,
	}

	err = db.Omit(clause.Associations).Create(&newProduct).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newProduct,
	})
}

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, _ := c.Params.Get("id")
	var product models.Product

	err := db.Where("id = ?", id).First(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, _ := c.Params.Get("id")
	var category models.Category

	err := db.Where("id = ?", id).First(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, _ := c.Params.Get("id")

	var product models.Product

	err := db.Where("id = ?", id).First(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var input ProductInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var categories []models.Category

	db.Find(&categories, input.Categories)

	updatedProduct := models.Product{
		Name:       input.Name,
		Categories: categories,
	}

	err = db.Model(&product).Updates(&updatedProduct).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, _ := c.Params.Get("id")

	var category models.Category

	err := db.Where("id = ?", id).First(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var input CategoryInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	updatedCategory := models.Product{
		Name: input.Name,
	}

	err = db.Model(&category).Updates(&updatedCategory).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": updatedCategory,
	})
}
