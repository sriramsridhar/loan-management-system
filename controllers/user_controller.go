package controllers

import (
	"github.com/gin-gonic/gin"
	"loan-management-system/models"
	"loan-management-system/services"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse("Users retrieved successfully", users))
}
