package routes

import (
	"psaa-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadKamar(e *gin.Engine) {
	e.GET("/kamar", controllers.ListKamar)
	e.GET("/kamar/:id", controllers.FindKamarByID)
	e.POST("/kamar", controllers.CreateKamar)
	e.PUT("/kamar/:id", controllers.UpdateKamar)
	e.DELETE("/kamar/:id", controllers.DeleteKamar)
}
