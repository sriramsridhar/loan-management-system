# Loan Management API

An API which facilitates the working of a minimal loan management system

## API context

- There are two types of users:- Customer, Admin
- Customer is the client who makes a request for the loan
- listing users, and making loan request on behalf of customer
- Before accessing to these functions an agent has to get the approval by the admin of being an agent
- Admin is the highest authority who can approve or reject a loan request
- Admin, cutomers and agents can login(agent can login only if agent is approved by the admin)
- When an agent will signup a request will be sent to the admin for approvalxxxxxx
- A loan has entities such as principle, interest rate, months to repay, emi and status
- The value of interest rate would depend on the value of principlexxxxxxxxx
- The loan can have 3 kinds of status: Pending, Approved, Rejected, Defaulted, Completed

Customer create a loan:
Customer submit a loan request defining amount and term
example:
Request amount of 10.000 $ with term 3 - weekly repayments
the loan and scheduled repayments will have state PENDING
1) Admin approve the loan:
Admin change the pending loans to state APPROVED
1) Customer can view loan belong to him:
Add a policy check to make sure that the customers can view them own loan only.
1) Customer add a repayments:
Customer add a repayment with amount greater or equal to the scheduled repayment
The scheduled repayment change the status to PAID
If all the scheduled repayments connected to a loan are PAID automatically also the loan become PAID

## Journeys

Admin - Signup(not required), Login, Approve/Reject loan requests, View all loans, View all customers, View Pending approval loans
Customer - Signup, Login, Create a loan, Repay a loan, View my loans, View my repayments

## Requirements 

- Go 1.22.4
- Postgres

## Key modules used

- Gin for the web framework.
- GORM for ORM.
- JWT for authentication.
- Viper for configuration management.

## How to run

Prometheus metrics exposed here : <http://localhost:8080/metrics>
