package repository

import (
	"backend-engineering-challenge/internals/domain/entity"
	"context"
)

type AccountRepositoryInterface interface {
	GetAccountDetailsByID(ctx context.Context, id string) ([]byte, error)
	GetAllAccountDetails(ctx context.Context) ([][]byte, error)
	DoTransaction(ctx context.Context, data entity.TransactionRequestEntity) (interface{}, error)
}
