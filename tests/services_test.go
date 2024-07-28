package services

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	testAdminUsername = "admin"
	testAdminPassword = "password"
)

// Initialize the test database
func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test_loan_management.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Customer{}, &Loan{}, &Repayment{})
	return db
}

// Setup the router for testing
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	authMiddleware := BasicAuthMiddleware{Username: testAdminUsername, Password: testAdminPassword}

	r.PUT("/loan/:id/approve", authMiddleware.Middleware(), func(c *gin.Context) {
		var loan Loan
		id := c.Param("id")

		if err := db.First(&loan, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
			return
		}

		adminUserID := c.MustGet("admin_user_id").(uint)

		loan.Approved = true
		loan.ApprovedUserID = adminUserID
		db.Save(&loan)

		c.JSON(http.StatusOK, loan)
	})

	return r
}

// Test the loan approval endpoint with authentication
func TestLoanApproval(t *testing.T) {
	db := InitTestDB()
	r := SetupRouter(db)

	// Create a customer and a loan
	customer := Customer{Name: "John Doe"}
	db.Create(&customer)

	loan := Loan{
		CustomerID: customer.ID,
		Amount:     10000,
		Term:       3,
		StartDate:  time.Now(),
	}
	db.Create(&loan)

	// Prepare the request
	req, _ := http.NewRequest("PUT", "/loan/1/approve", nil)
	req.SetBasicAuth(testAdminUsername, testAdminPassword)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the loan is approved
	var updatedLoan Loan
	db.First(&updatedLoan, loan.ID)
	assert.True(t, updatedLoan.Approved)
	assert.Equal(t, uint(1), updatedLoan.ApprovedUserID)
}

// Test the loan approval endpoint without authentication
func TestLoanApprovalWithoutAuth(t *testing.T) {
	db := InitTestDB()
	r := SetupRouter(db)

	// Create a customer and a loan
	customer := Customer{Name: "John Doe"}
	db.Create(&customer)

	loan := Loan{
		CustomerID: customer.ID,
		Amount:     10000,
		Term:       3,
		StartDate:  time.Now(),
	}
	db.Create(&loan)

	// Prepare the request without authentication
	req, _ := http.NewRequest("PUT", "/loan/1/approve", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check if the loan is not approved
	var updatedLoan Loan
	db.First(&updatedLoan, loan.ID)
	assert.False(t, updatedLoan.Approved)
}

func TestLoanApprovalWithWrongAuth(t *testing.T) {
	db := InitTestDB()
	r := SetupRouter(db)

	// Create a customer and a loan
	customer := Customer{Name: "John Doe"}
	db.Create(&customer)

	loan := Loan{
		CustomerID: customer.ID,
		Amount:     10000,
		Term:       3,
		StartDate:  time.Now(),
	}
	db.Create(&loan)

	// Prepare the request with wrong authentication
	req, _ := http.NewRequest("PUT", "/loan/1/approve", nil)
	req.SetBasicAuth("wrongUser", "wrongPass")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check if the loan is not approved
	var updatedLoan Loan
	db.First(&updatedLoan, loan.ID)
	assert.False(t, updatedLoan.Approved)
}
