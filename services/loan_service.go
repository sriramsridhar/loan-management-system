package services

import (
	"errors"
	"loan-management-system/database"
	"loan-management-system/models"
	"time"
)

func CreateLoan(loanReq *models.LoanReq) (models.Loan, error) {
	// You can add more logic here to calculate EMI based on the principal, interest rate, and months to repay
	startDate, err := time.Parse("2006-01-02", loanReq.StartDate)
	if err != nil {
		return models.Loan{}, errors.New("start date not found")
	}
	loan := models.Loan{
		CustomerID: loanReq.CustomerID,
		Amount:     loanReq.Amount,
		Term:       loanReq.Term,
		StartDate:  startDate,
		State:      models.LoanStatePending,
	}

	repaymentAmount := loanReq.Amount / float64(loanReq.Term)
	for i := 0; i < loanReq.Term; i++ {
		dueDate := startDate.AddDate(0, 0, 7*(i+1))
		if i == loanReq.Term-1 {
			repaymentAmount = loanReq.Amount - repaymentAmount*float64(loanReq.Term-1)
		}
		repayment := models.Repayment{
			Amount:  repaymentAmount,
			DueDate: dueDate,
			State:   models.Unpaid,
		}
		loan.Repayments = append(loan.Repayments, repayment)
	}
	database.DB.Create(&loan)
	return loan, nil
}

func GetLoans(user_id uint, user_role string) ([]models.Loan, error) {
	var loans []models.Loan
	if user_role == "admin" {
		err := database.DB.Preload("Repayments").Find(&loans).Error
		return loans, err
	}
	err := database.DB.Preload("Repayments").Where("customer_id = ?", user_id).Find(&loans).Error
	return loans, err
}

func GetLoanByID(id uint, user_id uint, user_role string) (models.Loan, error) {
	var loan models.Loan
	if user_role == "admin" {
		err := database.DB.Preload("Repayments").First(&loan, id).Error
		return loan, err
	}
	err := database.DB.Preload("Repayments").Where("customer_id = ?", user_id).First(&loan, id).Error
	return loan, err
}

func GetPendingApprovalLoans() ([]uint, error) {
	var loans []models.Loan
	var loanIds []uint
	err := database.DB.Preload("Repayments").Where("state = ?", models.LoanStatePending).Find(&loans).Error
	for i := range loans {
		loanIds = append(loanIds, loans[i].ID)
	}
	return loanIds, err
}

func ModifyLoanByID(id uint, adminId uint, state models.STATE) (models.Loan, error) {
	var loan models.Loan
	err := database.DB.Preload("Repayments").First(&loan, id).Error
	if err != nil {
		return loan, err
	}
	loan.State = state
	loan.ApprovedBy = adminId
	database.DB.Save(&loan)
	return loan, nil
}

func RepayLoan(loanId uint, amount float64) (models.Repayment, error) {
	var loan models.Loan
	var nextRepayment models.Repayment
	if err := database.DB.Preload("Repayments").First(&loan, loanId).Error; err != nil {
		return models.Repayment{}, errors.New("loan not found")
	}

	if loan.State != models.LoanStateApproved {
		return models.Repayment{}, errors.New("loan not approved")
	}

	for _, repayment := range loan.Repayments {
		if repayment.State != models.Paid && amount >= repayment.Amount {
			repayment.State = models.Paid
			amount -= repayment.Amount
			database.DB.Save(&repayment)
		} else if repayment.State != models.Paid && amount < repayment.Amount {
			nextRepayment = repayment
			break
		}
	}
	LoanPaid := checkAllPaid(loan)
	if LoanPaid {
		return models.Repayment{}, errors.New("loan completed")
	}
	return nextRepayment, nil
}

func checkAllPaid(loan models.Loan) bool {
	allPaid := true
	for _, r := range loan.Repayments {
		if r.State != models.Paid {
			allPaid = false
			break
		}
	}
	if allPaid {
		loan.State = models.LoanStateCompleted
		database.DB.Save(&loan)
	}
	return allPaid
}
