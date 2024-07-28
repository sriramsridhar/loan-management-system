package controllers

import (
	"loan-management-system/models"
	"loan-management-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.CreatedResponse("Signup successful", nil))
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	token, err := services.LoginUser(&user)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, models.NotFoundResponse(err.Error()))
		} else if err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse(err.Error()))
		} else if err.Error() == "agent not approved" {
			c.JSON(http.StatusForbidden, models.ForbiddenResponse(err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Login successful", gin.H{"token": token}))
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, models.SuccessResponse("Welcome to the loan management system", nil))
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, models.SuccessResponse("UP", nil))
}
