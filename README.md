# Loan Management API

An API which facilitates the working of a minimal loan management system

## API context

- There are two types of users:- Customer, Admin
- Customer is the client who makes a request for the loan
- Admin can listing users, and making loan request on behalf of customer
- Admin is the highest authority who can approve or reject a loan request
- The loan can have 3 kinds of status: Pending, Approved, Rejected, Defaulted, Completed

### Customer create a loan

    Customer submit a loan request defining amount and term
    the loan and scheduled repayments will have state PENDING

### Admin approve the loan

    Admin change the pending loans to state APPROVED

### Customer can view loan belong to him

    Add a policy check to make sure that the customers can view them own loan only.

### Customer add a repayments

    Customer add a repayment with amount greater or equal to the scheduled repayment
    The scheduled repayment change the status to PAID
    If all the scheduled repayments connected to a loan are PAID automatically also the loan become PAID

## Journeys

- Admin - Signup(not required), Login, Approve/Reject loan requests, View all loans, View all customers, View Pending approval loans
- Customer - Signup, Login, Create a loan, Repay a loan, View my loans, View my repayments

## Requirements

- Go 1.22.4
- Postgres

## Key modules used

- Gin for the web framework.
- GORM for ORM.
- JWT for authentication.
- Viper for configuration management.

## Container Deployment

Uses docker Compose for deployment
    ```bash
    make docker-build
    make docker-up
    ```

## Local development

make sure you have go installed in your system

    go mod init
    go mod tidy
    make run

I've used air in my local environment for hot reloading. you can install it using the following commands

    go install github.com/air-verse/air@latest
    air init
    air 

## Metrics & health check

- Health check exposed endpoint : <http://localhost:8080/health>
- Prometheus metrics exposed here : <http://localhost:8080/metrics>
