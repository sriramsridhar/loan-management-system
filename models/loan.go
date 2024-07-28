package models

import (
	"time"
)

type STATE string
type PAYMENTSTATE string

const (
	Paid   PAYMENTSTATE = "paid"
	Unpaid PAYMENTSTATE = "unpaid"
	Due    PAYMENTSTATE = "due"
)

const (
	LoanStateActive    STATE = "active"
	LoanStateApproved  STATE = "approved"
	LoanStateRejected  STATE = "rejected"
	LoanStateCompleted STATE = "completed"
	LoanStateDefaulted STATE = "defaulted"
	LoanStatePending   STATE = "pending"
)

type Customer struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Loans []Loan `gorm:"foreignKey:CustomerID"`
}

type Loan struct {
	ID         uint        `gorm:"primaryKey"`
	CustomerID uint        `json:"customer_id"`
	Amount     float64     `json:"amount"`
	Term       int         `json:"term"`
	State      STATE       `json:"state"`
	StartDate  time.Time   `json:"start_date"`
	ApprovedBy uint        `json:"approved_by"`
	Repayments []Repayment `gorm:"foreignKey:LoanID"`
}

type LoanReq struct {
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Term       int     `json:"term"`
	StartDate  string  `json:"start_date"` // in format "2006-01-02"
}

type Repayment struct {
	ID      uint         `gorm:"primaryKey"`
	LoanID  uint         `json:"loan_id"`
	Amount  float64      `json:"amount"`
	DueDate time.Time    `json:"due_date"`
	State   PAYMENTSTATE `json:"state"`
}

type RepaymentRequest struct {
	LoanID uint    `json:"loan_id"`
	Amount float64 `json:"amount"`
}
