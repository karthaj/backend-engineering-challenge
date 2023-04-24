package service

import (
	"backend-engineering-challenge/internals/app/repository"
	"backend-engineering-challenge/internals/app/usecase"
	req_res "backend-engineering-challenge/internals/domain/req-res"
	"context"
)

const usecasePrefixAccount = "backend-engineering-challenge.internals.app.usecase.ping"

type AccountServiceInterface interface {
	GetAccountDetailsByID(ctx context.Context, req req_res.GetAccountDetailsByIDRequest) (interface{}, error)
	GetAccountDetailsByName(ctx context.Context, req req_res.GetAccountDetailsByNameRequest) (interface{}, error)
	DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (interface{}, error)
}

type AccountService struct{}

func (r AccountService) GetAccountDetailsByID(ctx context.Context, req req_res.GetAccountDetailsByIDRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository)

	id := req.ID

	return data.GetAccountDetailsByID(ctx, id)
}
func (r AccountService) GetAccountDetailsByName(ctx context.Context, req req_res.GetAccountDetailsByNameRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository)

	id := req.Name

	return data.GetAccountDetailsByName(ctx, id)
}
func (r AccountService) DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (interface{}, error) {
	data := usecase.NewAccountUsecase(repository.AccountRepository)
	return data.DoTransaction(ctx, req)
}
