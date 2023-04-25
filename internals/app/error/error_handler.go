package errors

import (
	"backend-engineering-challenge/internals/domain"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// ErrorResponseValidation ...
type ErrorResponseValidation struct {
	Errors struct {
		Message       string      `json:"message"`
		Code          string      `json:"code"`
		CorrelationId string      `json:"correlationId"`
		Fields        interface{} `json:"fields"`
	} `json:"errors"`
}

// ErrorResponseDomain ...
type ErrorResponseDomain struct {
	Errors struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Line    int    `json:"line,omitempty"`
		File    string `json:"file,omitempty"`
		Trace   string `json:"trace,omitempty"`
	} `json:"errors"`
}

// ErrorResponseApplication ...
type ErrorResponseApplication struct {
	Errors struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Line    int    `json:"line,omitempty"`
		File    string `json:"file,omitempty"`
		Trace   string `json:"trace,omitempty"`
	} `json:"errors"`
}

// ErrorEncoder ...
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	switch reflect.TypeOf(err) {
	case reflect.TypeOf(AuthenticationError{}):
		encodeAuthenticationErrorResponse(ctx, err, w)

	case reflect.TypeOf(ApplicationError{}):
		encodeApplicationErrorResponse(ctx, err, w)

	case reflect.TypeOf(DomainError{}):
		encodeDomainErrorResponse(ctx, err, w)

	case reflect.TypeOf(ValidationError{}):
		encodeValidationErrorResponse(ctx, err, w)

	case reflect.TypeOf(GeneralError{}):
		encodeGeneralErrorResponse(ctx, err, w)

	}

}

type ApiErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	CorrelationId    string `json:"correlationId"`
	Code             string `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
}

func createErrorResponse(correlationId string, code string, message string, developerMessage string) ApiErrorResponse {
	response := ApiErrorResponse{}
	err := Error{
		CorrelationId:    correlationId,
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}

	response.Errors = []Error{err}
	return response
}

func encodeValidationErrorResponse(ctx context.Context, e error, w http.ResponseWriter) error {

	err := e.(ValidationError)
	errorCode := fmt.Sprintf("API-%+v", err.Code)

	correlationId := ctx.Value(domain.CorrelationIdContextKey).(string)
	if len(err.Message) != 0 {
		err.Message = "Validation Error"
	}

	res := ErrorResponseValidation{}
	res.Errors.Message = "Validation Error"
	res.Errors.Fields = err.Fields
	res.Errors.Code = errorCode
	res.Errors.CorrelationId = correlationId

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	return json.NewEncoder(w).Encode(res)

}

func encodeGeneralErrorResponse(_ context.Context, e error, w http.ResponseWriter) error {

	err := e.(GeneralError)

	res := createErrorResponse(err.CorrelationId, "API-"+strconv.Itoa(err.Code), err.Message, err.DeveloperMessage)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(res)

}

func encodeDomainErrorResponse(ctx context.Context, e error, w http.ResponseWriter) error {
	err := e.(DomainError)
	errorCode := fmt.Sprintf("API-%+v", err.Code)

	correlationId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res := createErrorResponse(correlationId, errorCode, err.Message, err.Trace)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(res)

}

func encodeApplicationErrorResponse(ctx context.Context, e error, w http.ResponseWriter) error {
	err := e.(ApplicationError)

	errorCode := fmt.Sprintf("API-%+v", err.Code)
	correlationId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res := createErrorResponse(correlationId, errorCode, err.Message, err.Trace)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(w).Encode(res)

}

func encodeAuthenticationErrorResponse(ctx context.Context, e error, w http.ResponseWriter) error {
	err := e.(AuthenticationError)
	errorCode := fmt.Sprintf("API-%+v", err.Code)
	correlationId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res := createErrorResponse(correlationId, errorCode, err.Message, err.Trace)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	return json.NewEncoder(w).Encode(res)
}
