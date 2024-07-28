# Loan Management API

An API which facilitates the working of a minimal loan management system

## API context

There are three types of users:- Customer, Agent, Admin
Customer is the client who makes a request for the loan
Agent is the middleman associated with the bank who has certain authourities such as edit loans,
listing users, and making loan request on behalf of customer
Before accessing to these functions an agent has to get the approval by the admin of being an agent
Admin is the highest authority who can approve or reject a loan and also the request by agent
Customers and agents can Sign up
Admin, cutomers and agents can login(agent can login only if agent is approved by the admin)
When an agent will signup a request will be sent to the admin for approvalxxxxxx
A loan has entities such as principle, interest rate, months to repay, emi and status
The value of interest rate would depend on the value of principlexxxxxxxxx
The loan can have 3 kinds of status: Approved, New or Rejected

## Tech stack used

Golang, Gin framework, Postgres

## Key modules used

Gin for the web framework.
GORM for ORM.
JWT for authentication.
Viper for configuration management.
Validator for input validation.

## How to run

Prometheus metrics exposed here : <http://localhost:8080/metrics>
