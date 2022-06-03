package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	LoadHelloRoutes(r)
	LoadUserRoutes(r)
	LoadAnakPanti(r)
	LoadKamar(r)
	LoadProducts(r)
	LoadPengasuh(r)

	return r
}
