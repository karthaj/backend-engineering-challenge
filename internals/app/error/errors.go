package errors

import (
	"fmt"
)

type DomainError struct {
	Message          string `json:"message"`
	Code             int    `json:"code"`
	Trace            string `json:"trace"`
	CorrelationId    string `json:"correlationId"`
	DeveloperMessage string `json:"developerMessage"`
}

func (e DomainError) Error() string {
	return fmt.Sprintf("Error Occurred %d", e.Code)
}

type ApplicationError struct {
	Message          string `json:"message"`
	Code             int    `json:"code"`
	Trace            string `json:"trace"`
	CorrelationId    string `json:"correlationId"`
	DeveloperMessage string `json:"developerMessage"`
}

type AuthenticationError struct {
	Message          string `json:"message"`
	Code             int    `json:"code"`
	Trace            string `json:"trace"`
	DeveloperMessage string `json:"developerMessage"`
}

func (e AuthenticationError) Error() string {
	return "Authentication Error"
}

func (e ApplicationError) Error() string {
	return fmt.Sprintf("Error Occurred %d", e.Code)
}

type ValidationError struct {
	Message          string      `json:"message"`
	CorrelationId    string      `json:"correlationId"`
	Fields           interface{} `json:"fields"`
	DeveloperMessage string      `json:"developerMessage"`
	Code             int         `json:"code"`
	Trace            string      `json:"trace"`
}

type GeneralError struct {
	Code             int         `json:"code"`
	CorrelationId    string      `json:"correlationId"`
	Message          string      `json:"message"`
	DeveloperMessage string      `json:"developerMessage"`
	Fields           interface{} `json:"fields"`
}

func (e ValidationError) Error() string {
	return "Validation Error"
}

func (e GeneralError) Error() string {
	return "General Error"
}

// NewAuthenticationError creates a new instance of AuthenticationError with the provided message and error code.
// Parameters:
// - messages: the error message.
// - code: the error code.
// Returns: a new instance of AuthenticationError.
func NewAuthenticationError(messages string, code int) AuthenticationError {
	return AuthenticationError{
		Message: messages,
		Code:    code,
	}
}

// NewDomainError creates a new instance of DomainError with the provided correlation ID, message, developer message and error code.
// Parameters:
// - correlationId: the correlation ID for the error.
// - message: the error message.
// - developerMessage: the error message for developers.
// - code: the error code.
// Returns: a new instance of DomainError.
func NewDomainError(correlationId string, message string, developerMessage string, code int) DomainError {
	return DomainError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

// NewApplicationError creates a new instance of ApplicationError with the provided correlation ID, message, developer message and error code.
// Parameters:
// - correlationId: the correlation ID for the error.
// - message: the error message.
// - developerMessage: the error message for developers.
// - code: the error code.
// Returns: a new instance of ApplicationError.
func NewApplicationError(correlationId string, message string, developerMessage string, code int) ApplicationError {
	return ApplicationError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

// NewValidationError creates a new instance of ValidationError with the provided correlation ID, message, developer message, fields and error code.
// Parameters:
// - correlationId: the correlation ID for the error.
// - message: the error message.
// - developerMessage: the error message for developers.
// - fields: a map of field names to error messages.
// - code: the error code.
// Returns: a new instance of ValidationError.
func NewValidationError(correlationId string, message string, developerMessage string, fields map[string]string, code int) ValidationError {
	return ValidationError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
		Fields:           fields,
	}
}
