package usecase

import (
	"backend-engineering-challenge/internals/domain/log"
	"backend-engineering-challenge/internals/domain/repository"
	"backend-engineering-challenge/internals/domain/req-res"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const usecasePrefixAccount = "backend-engineering-challenge.internals.app.usecase.account"

type AccountUsecase struct {
	AccountRepository repository.AccountRepositoryInterface
}

func NewAccountUsecase(ari repository.AccountRepositoryInterface) AccountUsecase {
	return AccountUsecase{
		ari,
	}
}

func (a AccountUsecase) GetAccountDetailsByID(ctx context.Context, id string) (req_res.AccountResponse, error) {

	temp := req_res.AccountEntity{}

	data, err := a.AccountRepository.GetAccountDetailsByID(ctx, id)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to get value from DB")
	}

	err = json.Unmarshal(data, &temp)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to parse value to account object")
	}

	var res req_res.AccountResponse
	res.ID = temp.ID
	res.Name = temp.Name
	res.Balance, _ = strconv.ParseFloat(temp.Balance, 64)

	return res, nil
}

func (a AccountUsecase) GetAccountDetailsByName(ctx context.Context, name string) (req_res.AccountResponse, error) {

	res := req_res.AccountResponse{}

	data, err := a.AccountRepository.GetAccountDetailsByName(ctx, name)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to get value from DB")
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to parse value to account object")
	}

	return res, nil
}

func (a AccountUsecase) DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (req_res.AccountResponse, error) {
	res := req_res.AccountResponse{}

	data, err := a.AccountRepository.GetAccountDetailsByID(ctx, req.ID)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to get value from DB")
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.ErrorContext(
			ctx,
			usecasePrefixAccount,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to parse value to account object")
	}

	return res, nil
}
