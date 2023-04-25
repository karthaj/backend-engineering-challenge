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

// GetAllAccountDetails returns all the account details.
// If no account is found, returns an error
// Otherwise, returns the details of the accounts as an []AccountResponse.
// If an error occurs while retrieving the account details, returns an error response.
func (a AccountUsecase) GetAllAccountDetails(ctx context.Context) ([]req_res.AccountResponse, error) {

	corId := ctx.Value(domain.CorrelationIdContextKey).(string)

	res := []req_res.AccountResponse{}

	data, err := a.AccountRepository.GetAllAccountDetails(ctx)
	if data == nil {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "No accounts")

		return res, errors.NewDomainError(corId, "No accounts available", "", errors.ErrAccountNotFound)
	}

	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to get value from DB")

		return res, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	for _, datum := range data {
		var t entity.AccountEntity
		err = json.Unmarshal(datum, &t)
		if err != nil {
			log.ErrorContext(
				ctx,
				fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "Unable to parse value to account object")

			return res, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead)
		}
		tt, _ := strconv.ParseFloat(t.Balance, 64)
		res = append(res, req_res.AccountResponse{
			ID: t.ID, Name: t.Name, Balance: tt,
		})
	}

	if len(res) == 0 {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByName"), "No accounts found")

		return res, errors.NewApplicationError(corId, "No Accounts found", "", errors.ErrBatabaseRead)
	}

	return res, nil
}

// DoTransaction performs an account transaction operation, ensuring thread-safety using a mutex
// It takes a context.Context object and a req_res.DoTransactionRequest object as input and returns a
// req_res.TransactionAccResponse and error. It acquires a mutex lock before executing the transaction to ensure
// thread-safety and releases the lock once the transaction is complete.
func (a AccountUsecase) DoTransaction(ctx context.Context, req req_res.DoTransactionRequest) (req_res.TransactionAccResponse, error) {

	a.M.Lock()         // acquire lock on account
	defer a.M.Unlock() // release lock when function returns

	corId := ctx.Value(domain.CorrelationIdContextKey).(string) // get correlation ID from context

	if req.FromAccountId == req.ToAccountId { // check if the transaction is being made to the same account
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "You cannot do transactions to the same account", fmt.Sprintf("ID : %v", req.ToAccountId))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "You cannot do transactions to the same account", "", errors.ErrAccountNotFound) // return error if transaction is being made to the same account
	}

	fromAcc := entity.AccountEntity{} // create empty AccountEntity objects for sender and receiver accounts
	toAcc := entity.AccountEntity{}

	//  <<<<<<<<<< FROM ACCOUNT
	//  <<<<<<<<<< FROM ACCOUNT
	//  <<<<<<<<<< FROM ACCOUNT
	fromAccData, err := a.AccountRepository.GetAccountDetailsByID(ctx, req.FromAccountId) // get account details for sender
	// check if sender account exists
	if fromAccData == nil {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Account not found", fmt.Sprintf("FromAccountId ID : %v", req.FromAccountId))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Invalid sender account to perform transaction", "", errors.ErrAccountNotFound) // return error if sender account is invalid
	}

	if err != nil { // check for any errors while retrieving account details
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to get value from DB")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead) // return application error if there was an error while retrieving account details
	}

	err = json.Unmarshal(fromAccData, &fromAcc) // unmarshal account details to sender account object
	// check for any errors while unmarshalling account details
	if err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to parse value to account object")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead) // return application error if there was an error while unmarshaling account details
	}

	if fromAcc.ID == "" { // check if sender account ID is empty
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to query")

		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead) // return application error if there was an error while querying sender account details
	}

	balance, _ := strconv.ParseFloat(fromAcc.Balance, 64) // convert sender account balance to float64
	if balance < req.Amount {
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "DoTransaction"), "Insufficient account balance", fmt.Sprintf("req : %+v", req), fmt.Sprintf("account : %+v", toAcc))

		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Insufficient account balance", "", errors.ErrAccountBalanceInsufficient)
	}

	//  >>>>>>>>>>> TO ACCOUNT
	//  >>>>>>>>>>> TO ACCOUNT
	//  >>>>>>>>>>> TO ACCOUNT
	// Get account details of the sender
	toAccData, err := a.AccountRepository.GetAccountDetailsByID(ctx, req.ToAccountId)
	if toAccData == nil {
		// Log account not found error
		log.DebugContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Accounts not found", fmt.Sprintf("ToAccountId ID : %v", req.ToAccountId))
		// Return domain error for invalid account
		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Invalid receiver account to perform transaction", "", errors.ErrAccountNotFound)
	}

	if err != nil {
		// Log error for DB read failure
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to get value from DB")
		// Return application error for internal error
		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	// Unmarshal JSON data to account object
	err = json.Unmarshal(toAccData, &toAcc)
	if err != nil {
		// Log error for parsing data failure
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to parse value to account object")
		// Return application error for internal error
		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Unable to fetch data", "Unable to parse the value to the account object", errors.ErrBatabaseRead)
	}

	if toAcc.ID == "" {
		// Log error for DB query failure
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "GetAccountDetailsByID"), "Unable to query")
		// Return application error for internal error
		return req_res.TransactionAccResponse{}, errors.NewApplicationError(corId, "Internal error occurred", "DB error", errors.ErrBatabaseRead)
	}

	// Convert balance strings to float64
	fromBal, _ := strconv.ParseFloat(fromAcc.Balance, 64)
	toBal, _ := strconv.ParseFloat(toAcc.Balance, 64)

	// Deduct transaction amount from sender balance and add it to receiver balance
	fromAcc.Balance = fmt.Sprintf("%2f", fromBal-req.Amount)
	toAcc.Balance = fmt.Sprintf("%2f", toBal+req.Amount)

	// Create transaction request entity
	requestEntity := entity.TransactionRequestEntity{
		ToAcc:   toAcc,
		FromAcc: fromAcc,
	}

	// Perform transaction in repository
	_, err = a.AccountRepository.DoTransaction(ctx, requestEntity)
	if err != nil {
		// Log error for insufficient balance and failed transaction
		log.ErrorContext(
			ctx,
			fmt.Sprintf("%s.%s", usecasePrefixAccount, "DoTransaction"), "Transaction failed", fmt.Sprintf("req : %+v", req), fmt.Sprintf("requestEntity : %+v", requestEntity))
		// Return domain error for transaction failure
		return req_res.TransactionAccResponse{}, errors.NewDomainError(corId, "Transaction failed", "", errors.ErrTransactionFailed)
	}

	// Create transaction response object
	res := req_res.TransactionAccResponse{
		ID:        req.ToAccountId,
		Name:      toAcc.Name,
		Balance:   toBal + req.Amount,
		Reference: time.Now().Format("TRX-20060201150405"),
	}

	return res, nil
}
