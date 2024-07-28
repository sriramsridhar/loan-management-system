package services

import (
	"loan-management-system/database"
	"loan-management-system/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}
