package routes

import (
	"loan-management-system/controllers"
	"loan-management-system/middlewares"

	"github.com/gin-gonic/gin"
)

func LoanRoutes(router *gin.Engine) {
	loan := router.Group("/loan").Use(middlewares.AuthMiddleware())
	{
		loan.POST("/", controllers.CreateLoan)
		loan.GET("/getall", controllers.GetLoans)
		loan.GET("/:id", controllers.GetLoanByID)
		loan.GET("/pending_approval", controllers.GetPendingApprovalLoans)
		loan.GET("/approve/:id", controllers.ApproveLoan)
		loan.GET("/reject/:id", controllers.RejectLoan)
		loan.POST("/repay", controllers.RepayLoan)
	}
}
