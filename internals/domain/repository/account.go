package repository

import (
	"backend-engineering-challenge/internals/domain/entity"
	"context"
)

type AccountRepositoryInterface interface {
	GetAccountDetailsByID(ctx context.Context, id string) ([]byte, error)
	GetAccountDetailsByName(ctx context.Context, name string) ([]byte, error)
	DoTransaction(ctx context.Context, data entity.AccountEntity) (interface{}, error)
}
