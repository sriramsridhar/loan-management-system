// main.go
package main

import (
	"loan-management-system/config"
	"loan-management-system/database"
	"loan-management-system/routes"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// @title Loan Management System API
// @version 1.0
// @description This is a sample server for managing loans.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func main() {
	config.LoadConfig()
	database.Connect()
	defer database.DB.Close()
	if config.AppConfig.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Migrate the schema
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	routes.AuthRoutes(r)
	routes.LoanRoutes(r)
	routes.UserRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}
