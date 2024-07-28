package routes

import (
	"loan-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}

	home := router.Group("/")
	{
		home.GET("/home", controllers.Home)
		home.GET("/health", controllers.Health)
	}

}
