package req_res

// AccountResponse struct represents the account information.
type AccountResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// AccountResponse struct represents the account information.
type AccountEntity struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type DoTransactionRequest struct {
	ID              string  `json:"id"  validate:"required"`
	Amount          float64 `json:"amount"  validate:"required"`
	TransactionType string  `json:"transactionType"  validate:"required"`
}

type GetAccountDetailsByIDRequest struct {
	ID string `json:"id"  validate:"required"`
}

type GetAccountDetailsByNameRequest struct {
	Name string `json:"name"  validate:"required"`
}

type DoTransactionResponse struct {
	Data struct {
		Account AccountResponse `json:"account"`
	} `json:"data"`
	Meta CommonMetaResponse `json:"meta"`
}
