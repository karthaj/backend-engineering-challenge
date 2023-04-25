package codec

import (
	errors "backend-engineering-challenge/internals/app/error"
	"backend-engineering-challenge/internals/domain"
	"backend-engineering-challenge/internals/domain/log"
	"backend-engineering-challenge/internals/domain/req-res"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"runtime/debug"
	"strings"
)

const logPrefixAccountTransaction = `backend-engineering-challenge.internals.transport.http.codec.account`

func DecodeDoTransaction(ctx context.Context, req *http.Request) (request interface{}, err error) {
	var requestData req_res.DoTransactionRequest
	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccountTransaction, "DecodeDoTransaction"),
			fmt.Sprintf(`Decode DoTransactionRequest, json decode, Error:%v`, err))

		errs := make(map[string]string)
		errs["json"] = err.Error()

		return nil, errors.NewValidationError(corId, "Invalid request", string(debug.Stack()), errs, errors.ErrRequestInvalid)
	}

	validate := validator.New()
	err = validate.Struct(&requestData)

	if err != nil {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccountTransaction, "DecodeDoTransaction"),
			fmt.Sprintf(`Parse DoTransactionRequest, json parse, Error:%v`, err.Error()))

		errs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errs[strings.Split(err.StructNamespace(), ".")[1]] = err.Tag()
		}
		return nil, errors.NewValidationError(corId, "Invalid request", string(debug.Stack()), errs, errors.ErrRequestInvalid)

	}

	return requestData, nil

}

func EncodeDoTransaction(ctx context.Context, w http.ResponseWriter, data interface{}) (err error) {
	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res, ok := data.(req_res.TransactionAccResponse)
	if !ok {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccountTransaction, "DecodeDoTransaction"),
			fmt.Sprintf(`Parse DoTransactionRequest, json parse, Error:%v`, err.Error()))

		return errors.NewDomainError(corId, "Unable to process response", string(debug.Stack()), errors.ErrResponseParse)
	}

	result := req_res.DoTransactionResponse{}
	result.Meta.Code = 200
	result.Meta.Message = "Success"
	result.Data.Account = res
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(result)

}
