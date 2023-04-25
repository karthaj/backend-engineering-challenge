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
	Amount        float64 `json:"amount" validate:"required"`
	ToAccountId   string  `json:"toAccountId" validate:"required"`
	FromAccountId string  `json:"fromAccountId" validate:"required"`
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
		Accounts []AccountResponse `json:"accounts"`
	} `json:"data"`
	Meta CommonMetaResponse `json:"meta"`
}
