package req_res

// AccountResponse struct represents the account information.
type AccountResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// TransactionAccResponse struct represents the account information.
type TransactionAccResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	Reference string  `json:"reference"`
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
		Account TransactionAccResponse `json:"account"`
	} `json:"data"`
	Meta CommonMetaResponse `json:"meta"`
}

type GeneralAccountResponse struct {
	Data struct {
		Account AccountResponse `json:"reference"`
	} `json:"data"`
	Meta CommonMetaResponse `json:"meta"`
}
