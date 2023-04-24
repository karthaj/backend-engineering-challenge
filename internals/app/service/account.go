package service

import (
	"backend-engineering-challenge/internals/app/repository"
	"backend-engineering-challenge/internals/app/usecase"
	req_res "backend-engineering-challenge/internals/domain/req-res"
	"context"
	"sync"
)

// AccountServiceInterface defines the methods that should be implemented by any
// struct that provides account-related services.
type AccountServiceInterface interface {

	// GetAccountDetailsByID retrieves the account details by the account ID.
	// The returned response is of type req_res.AccountResponse, which contains
	// details such as the account ID, name, balance, and transaction history.
	// The function returns an error if there was a problem with the request.
	GetAccountDetailsByID(ctx context.Context, req req_res.GetAccountDetailsByIDRequest) (interface{}, error)

	// GetAccountDetailsByName retrieves the account details by the account name.
	// The returned response is of type req_res.AccountResponse, which contains
	// details such as the account ID, name, balance, and transaction history.
	// The function returns an error if there was a problem with the request.
	GetAccountDetailsByName(ctx context.Context, req req_res.GetAccountDetailsByNameRequest) (interface{}, error)

	// DoTransaction performs a transaction between two accounts. The request
	// should contain details such as the sender and recipient account IDs, and
	// the transaction amount. The function returns a response of type
	// req_res.TransactionAccResponse, which contains details such as the
	// transaction status and the updated account balances of the sender and
	// recipient. The function returns an error if there was a problem with
	// the request.
	DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (interface{}, error)
}

type AccountService struct{}

var M = &sync.Mutex{}

func (r AccountService) GetAccountDetailsByID(ctx context.Context, req req_res.GetAccountDetailsByIDRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository, M)
	id := req.ID
	return data.GetAccountDetailsByID(ctx, id)
}
func (r AccountService) GetAccountDetailsByName(ctx context.Context, req req_res.GetAccountDetailsByNameRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository, M)
	id := req.Name
	return data.GetAccountDetailsByName(ctx, id)
}
func (r AccountService) DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository, M)
	return data.DoTransaction(ctx, req)
}
