package service

import (
	"backend-engineering-challenge/internals/app/usecase"
	"backend-engineering-challenge/internals/domain/log"
	"context"
	"fmt"
)

const usecasePrefixPing = "internals.app.usecase.ping"

type PingService interface {
	Pinging(ctx context.Context) (string, error)
}

type PingServiceStr struct{}

func (r PingServiceStr) Pinging(ctx context.Context) (string, error) {

	uc := usecase.NewPingUsecase()

	res, err := uc.Ping()

	if err != nil {
		log.ErrorContext(ctx, fmt.Sprint(usecasePrefixPing),
			"Pinging failed",
			fmt.Sprint("Pinging was successful"),
			fmt.Sprintf("%+v", res))
		return res, err

	}

	log.InfoContext(ctx, fmt.Sprint(usecasePrefixPing),
		"Pinging successfully",
		fmt.Sprint("Pinging was successful"),
		fmt.Sprintf("%+v", res))

	return res, nil

}
