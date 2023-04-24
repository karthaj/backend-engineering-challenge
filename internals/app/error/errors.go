package errors

import (
	"fmt"
	"reflect"
)

type DomainError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Trace   string `json:"trace"`
}

func (e DomainError) Error() string {
	return fmt.Sprintf("Error Occurred %d", e.Code)
}

type ApplicationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Trace   string `json:"trace"`
}

type AuthenticationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Trace   string `json:"trace"`
}

func (e AuthenticationError) Error() string {
	return "Authentication Error"
}

func (e ApplicationError) Error() string {
	return fmt.Sprintf("Error Occurred %d", e.Code)
}

type ValidationError struct {
	Message string      `json:"message"`
	Fields  interface{} `json:"fields"`
	Code    int         `json:"code"`
	Trace   string      `json:"trace"`
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

func NewAuthenticationError(messages string, code int) AuthenticationError {

	return AuthenticationError{
		Message: messages,
		Code:    code,
	}
}

func NewDomainError(correlationId string, message string, developerMessage string, code int) GeneralError {
	return GeneralError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func NewApplicationError(correlationId string, message string, developerMessage string, code int) GeneralError {
	return GeneralError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func NewValidationError(correlationId string, message string, developerMessage string, fields map[string]string, code int) GeneralError {
	return GeneralError{
		Code:             code,
		CorrelationId:    correlationId,
		Message:          message,
		DeveloperMessage: developerMessage,
		Fields:           fields,
	}
}

func IsValidationError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(GeneralError{})
}

func IsApplicationError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(GeneralError{})
}

func IsDomainError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(GeneralError{})
}