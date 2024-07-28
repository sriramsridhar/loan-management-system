package routes

import (
	"github.com/gin-gonic/gin"
	"loan-management-system/controllers"
	"loan-management-system/middlewares"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user").Use(middlewares.AuthMiddleware())
	{
		user.GET("/", controllers.GetUsers)
	}
}
