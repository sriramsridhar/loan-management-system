package database

import (
	"loan-management-system/config"
	"loan-management-system/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=" + config.AppConfig.DBHost + " port=" + config.AppConfig.DBPort + " user=" + config.AppConfig.DBUsername + " dbname=" + config.AppConfig.DBName + " password=" + config.AppConfig.DBPassword + " sslmode=disable"
	var err error
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Loan{}, &models.Repayment{})
}
