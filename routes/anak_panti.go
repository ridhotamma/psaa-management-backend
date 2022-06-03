package routes

import (
	"psaa-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadAnakPanti(e *gin.Engine) {
	e.POST("/anak-panti", controllers.CreateAnakPanti)
	e.GET("/anak-panti", controllers.AnakPantiList)
	e.GET("/anak-panti/:id", controllers.FindAnakPantiById)
	e.DELETE("/anak-panti/:id", controllers.DeleteAnakPanti)
	e.PATCH("/anak-panti/:id", controllers.UpdateAnakPanti)
}
