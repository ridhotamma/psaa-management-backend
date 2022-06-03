package routes

import (
	"psaa-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadPengasuh(e *gin.Engine) {
	e.POST("/pengasuh", controllers.CreatePengasuh)
}
