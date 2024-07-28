package models

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
}

func CreatedResponse(message string, data interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	}
}

func BadRequestResponse(message string) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func UnauthorizedResponse(message string) Response {
	return Response{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

func ForbiddenResponse(message string) Response {
	return Response{
		Status:  http.StatusForbidden,
		Message: message,
	}
}

func NotFoundResponse(message string) Response {
	return Response{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func InternalServerErrorResponse(message string) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}
