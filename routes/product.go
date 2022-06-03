package routes

import (
	"psaa-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadProducts(e *gin.Engine) {
	e.GET("/product", controllers.ListProduct)
	e.GET("/category", controllers.ListCategory)
	e.POST("/product", controllers.CreateProduct)
	e.POST("/category", controllers.CreateCategory)
}
