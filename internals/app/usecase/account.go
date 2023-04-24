package usecase

import (
	errors "backend-engineering-challenge/internals/app/error"
	"backend-engineering-challenge/internals/domain"
	"backend-engineering-challenge/internals/domain/entity"
	"backend-engineering-challenge/internals/domain/log"
	"backend-engineering-challenge/internals/domain/repository"
	"backend-engineering-challenge/internals/domain/req-res"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const usecasePrefixAccount = "backend-engineering-challenge.internals.app.usecase.account"

type AccountUsecase struct {
	AccountRepository repository.AccountRepositoryInterface
	M                 *sync.Mutex
}

func NewAccountUsecase(ari repository.AccountRepositoryInterface, m *sync.Mutex) AccountUsecase {
	return AccountUsecase{
		AccountRepository: ari,
		M:                 m,
	}
}

// GetAccountDetailsByID returns the account details with the given name.
// If no account is found with the given name, returns an error
// Otherwise, returns the details of the account as an AccountResponse.
// If an error occurs while retrieving the account details, returns an error response.
func (a AccountUsecase) GetAccountDetailsByID(ctx context.Context, id string) (req_res.AccountResponse, error) {

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	var temp entity.AccountEntity

	data, err := a.AccountRepository.GetAccountDetailsByID(ctx, id)
	if data == nil {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "No data found", fmt.Sprintf("ID : %v", id))

		return req_res.AccountResponse{}, errors.NewDomainError(corId, "Account not found", "", errors.ErrAccountNotFound)
	}

	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to get value from DB")

		return req_res.AccountResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	err = json.Unmarshal(data, &temp)
	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to parse value to account object")

		return req_res.AccountResponse{}, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead)
	}

	var res req_res.AccountResponse
	res.ID = temp.ID
	res.Name = temp.Name
	res.Balance, _ = strconv.ParseFloat(temp.Balance, 64)

	return res, nil
}

// GetAccountDetailsByName returns the account details with the given name.
// If no account is found with the given name, returns an error
// Otherwise, returns the details of the account as an AccountResponse.
// If an error occurs while retrieving the account details, returns an error response.
func (a AccountUsecase) GetAccountDetailsByName(ctx context.Context, name string) (req_res.AccountResponse, error) {

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res := req_res.AccountResponse{}

	data, err := a.AccountRepository.GetAccountDetailsByName(ctx, name)
	if data == nil {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "No data found", fmt.Sprintf("Name : %v", name))

		return res, errors.NewDomainError(corId, "Account not found", "", errors.ErrAccountNotFound)
	}

	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to get value from DB")

		return res, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to parse value to account object")

		return res, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead)
	}

	return res, nil
}

// DoTransaction performs an account transaction operation, ensuring thread-safety using a mutex
// It takes a context.Context object and a req_res.DoTransactionRequest object as input and returns a
// req_res.TransactionAccResponse and error. It acquires a mutex lock before executing the transaction to ensure
// thread-safety and releases the lock once the transaction is complete.
func (a AccountUsecase) DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (req_res.TransactionAccResponse, error) {

	a.M.Lock()
	defer a.M.Unlock()

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	account := entity.AccountEntity{}

	data, err := a.AccountRepository.GetAccountDetailsByID(ctx, req.ID)
	if data == nil {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Account not found", fmt.Sprintf("ID : %v", req.ID))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Invalid account to perform transaction", "", errors.ErrAccountNotFound)
	}

	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to get value from DB")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	err = json.Unmarshal(data, &account)
	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to parse value to account object")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead)
	}

	if account.ID == "" {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to query")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	balance, _ := strconv.ParseFloat(account.Balance, 64)

	if req.TransactionType == "CR" && balance <= req.Amount {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "DoTransaction"), "Insufficient account balance", fmt.Sprintf("req : %+v", req), fmt.Sprintf("account : %+v", account))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Insufficient account balance", "", errors.ErrAccountBalanceInsufficient)
	}

	if req.TransactionType == domain.TrxDebit {
		// credit the amount
		req.Amount = -1 * req.Amount
	}

	balance += req.Amount

	trx := entity.AccountEntity{
		ID:      req.ID,
		Name:    account.Name,
		Balance: fmt.Sprintf("%2f", balance),
	}

	_, err = a.AccountRepository.DoTransaction(ctx, trx)
	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "DoTransaction"), "Insufficient account balance", fmt.Sprintf("req : %+v", req), fmt.Sprintf("account : %+v", account))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Transaction failed", "", errors.ErrTransactionFailed)
	}

	res := req_res.TransactionAccResponse{
		ID:        req.ID,
		Name:      account.Name,
		Balance:   balance,
		Reference: time.Now().Format(req.TransactionType + "20060201150405"),
	}

	return res, nil
}
