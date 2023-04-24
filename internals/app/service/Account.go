package service

import (
	"backend-engineering-challenge/internals/domain/entity"
	"context"
)

type AccountService interface {
	GetAccountDetailsByID(ctx context.Context, id string) (entity.Account, error)
	GetAccountDetailsByName(ctx context.Context, name string) (entity.Account, error)
	DoTransaction(ctx context.Context, data entity.TransactionRequest) (entity.Account, error)
}
