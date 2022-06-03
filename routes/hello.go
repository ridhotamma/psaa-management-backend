package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadHelloRoutes(e *gin.Engine) {
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from routes hello",
		})
	})

	e.GET("/greeting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "greeting from routes hello",
		})
	})
}
