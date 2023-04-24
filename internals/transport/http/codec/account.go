package codec

import (
	errors "backend-engineering-challenge/internals/app/error"
	"backend-engineering-challenge/internals/domain"
	"backend-engineering-challenge/internals/domain/log"
	"backend-engineering-challenge/internals/domain/req-res"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"runtime/debug"
	"strings"
)

const logPrefixAccount = `backend-engineering-challenge.internals.transport.http.codec.account`

func DecodeGetAccountDetailsByID(ctx context.Context, req *http.Request) (request interface{}, err error) {

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)
	params := mux.Vars(req)
	id := strings.Trim(params["id"], " ")

	if id == "" {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccount, "GetAccountDetailsByIDRequest"),
			`DecodeGetAccountDetailsByID, ID not found`)

		errs := make(map[string]string)
		errs["id"] = "param | required"

		return nil, errors.NewValidationError(corId, "ID required", string(debug.Stack()), errs, errors.ErrParamNotFound)
	}

	var data req_res.GetAccountDetailsByIDRequest
	data.ID = id

	return data, nil

}

func EncodeGetAccountDetailsByID(ctx context.Context, w http.ResponseWriter, data interface{}) (err error) {
	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res, ok := data.(req_res.AccountResponse)
	if !ok {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccount, "DecodeDoTransaction"),
			fmt.Sprintf(`Parse DoTransactionRequest, json parse, Error:%v`, err.Error()))

		return errors.NewDomainError(corId, "Unable to process response", string(debug.Stack()), errors.ErrResponseParse)
	}

	result := req_res.GeneralAccountResponse{}
	result.Meta.Code = 200
	result.Meta.Message = "Success"
	result.Data.Account = res
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(result)

}

func DecodeGetAccountDetailsByName(ctx context.Context, req *http.Request) (request interface{}, err error) {

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)
	params := mux.Vars(req)
	name := strings.Trim(params["name"], " ")

	if name == "" {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccount, "GetAccountDetailsByIDRequest"),
			`DecodeGetAccountDetailsByName, Name not found`)

		errs := make(map[string]string)
		errs["name"] = "param | required"

		return nil, errors.NewValidationError(corId, "Name required", string(debug.Stack()), errs, errors.ErrParamNotFound)
	}

	var data req_res.GetAccountDetailsByNameRequest
	data.Name = name

	return data, nil

}

func EncodeGetAccountDetailsByName(ctx context.Context, w http.ResponseWriter, data interface{}) (err error) {
	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res, ok := data.(req_res.AccountResponse)
	if !ok {
		log.ErrorContext(
			ctx, fmt.Sprintf("%s.%s", logPrefixAccount, "DecodeDoTransaction"),
			fmt.Sprintf(`Parse DoTransactionRequest, json parse, Error:%v`, err.Error()))

		return errors.NewDomainError(corId, "Unable to process response", string(debug.Stack()), errors.ErrResponseParse)
	}

	result := req_res.GeneralAccountResponse{}
	result.Meta.Code = 200
	result.Meta.Message = "Success"
	result.Data.Account = res
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(result)

}
