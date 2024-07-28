package controllers

import (
	"loan-management-system/models"
	"loan-management-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanRequest struct {
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Term       int     `json:"term"`
	StartDate  string  `json:"start_date"` // in format "2006-01-02"
}

func getuserId(c *gin.Context) uint {
	userIdFloat := c.MustGet("user_id").(float64)
	userId := strconv.FormatFloat(userIdFloat, 'f', -1, 64)
	userIdInt, _ := strconv.ParseUint(userId, 10, 32)
	return uint(userIdInt)
}

func CreateLoan(c *gin.Context) {
	var req models.LoanReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// customer can create loan only for him and admin can create for anyone
	if c.MustGet("role") != "admin" {
		req.CustomerID = getuserId(c) // we can throw error if customer id is not same as user id but i overrided it for now.
	}
	loan, err := services.CreateLoan(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.CreatedResponse("Loan created successfully", loan))
}

func GetLoans(c *gin.Context) {
	user_id := getuserId(c)
	loans, err := services.GetLoans(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Loans retrieved successfully", loans))
}

func GetLoanByID(c *gin.Context) {
	id := c.Param("id")
	user_id := getuserId(c)
	loanID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Invalid loan ID"))
		return
	}

	loan, err := services.GetLoanByID(uint(loanID), user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NotFoundResponse("Loan not found"))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Loan retrieved successfully", loan))
}

func GetPendingApprovalLoans(c *gin.Context) {
	if c.MustGet("role") != "admin" {
		c.JSON(http.StatusForbidden, models.ForbiddenResponse("You are not authorized to view pending approval loans"))
		return
	}
	loanIds, err := services.GetPendingApprovalLoans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Pending approval loans retrieved successfully", loanIds))
}

func ApproveLoan(c *gin.Context) {
	id := c.Param("id")
	if c.MustGet("role") != "admin" {
		c.JSON(http.StatusForbidden, models.ForbiddenResponse("You are not authorized to approve loans"))
		return
	}
	adminId := getuserId(c)
	loanId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Invalid loan ID"))
		return
	}

	loan, _ := services.ModifyLoanByID(uint(loanId), adminId, models.LoanStateApproved)
	c.JSON(http.StatusOK, models.SuccessResponse("Loan approved successfully", loan))
}

func RejectLoan(c *gin.Context) {
	id := c.Param("id")
	if c.MustGet("role") != "admin" {
		c.JSON(http.StatusForbidden, models.ForbiddenResponse("You are not authorized to reject loans"))
		return
	}
	adminId := getuserId(c)
	loanId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Invalid loan ID"))
		return
	}

	loan, _ := services.ModifyLoanByID(uint(loanId), adminId, models.LoanStateRejected)
	c.JSON(http.StatusOK, models.SuccessResponse("Loan rejected successfully", loan))
}

func RepayLoan(c *gin.Context) {
	var req models.RepaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Invalid request format"))
		return
	}

	repayment, err := services.RepayLoan(req.LoanID, req.Amount)
	if err == nil {
		c.JSON(http.StatusOK, models.SuccessResponse("Loan repayment completed successfully", repayment))
	}
	if err.Error() == "loan not found" {
		c.JSON(http.StatusNotFound, models.NotFoundResponse("Loan not found"))
	}
	if err.Error() == "loan not approved" {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Loan not approved"))
	}
	if err.Error() == "loan completed" {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("Loan already completed"))
	}

}
