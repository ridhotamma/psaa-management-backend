package routes

import (
	"psaa-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(e *gin.Engine) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user", controllers.UserList)
	e.GET("/user/:id", controllers.FindUserById)
	e.DELETE("/user/:id", controllers.DeleteUser)
	e.PATCH("/user/:id", controllers.UpdateUser)
}
