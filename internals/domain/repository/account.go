package repository

import (
	"backend-engineering-challenge/internals/domain/entity"
	"backend-engineering-challenge/internals/domain/req-res"
	"context"
)

type AccountRepositoryInterface interface {
	GetAccountDetailsByID(ctx context.Context, id string) ([]byte, error)
	GetAccountDetailsByName(ctx context.Context, name string) ([]byte, error)
	DoTransaction(ctx context.Context, data req_res.DoTransactionRequest) (entity.Account, error)
}
