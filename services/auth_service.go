package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"loan-management-system/database"
	"loan-management-system/models"
	"loan-management-system/utils"
)

func CreateUser(user *models.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return database.DB.Create(user).Error
}

func LoginUser(user *models.User) (string, error) {
	var dbUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !dbUser.Approved && dbUser.Role == "agent" {
		return "", errors.New("agent not approved")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(dbUser.ID, dbUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
