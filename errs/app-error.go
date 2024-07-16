package errs

import "net/http"

// AppError ...
type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// AsMessage ...
func (e *AppError) AsMessage() *AppError {

	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError ...
func NewNotFoundError(message string) *AppError {

	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}

}

// NewUnexpectedError ...
func NewUnexpectedError(message string) *AppError {

	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}

}

// NewValidationError ...
func NewValidationError(message string) *AppError {

	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}

}

// NewBadRequestError ...
func NewBadRequestError(message string) *AppError {

	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}

}
