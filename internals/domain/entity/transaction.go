package entity

type TransactionRequestEntity struct {
	ToAcc   AccountEntity `json:"ToAcc"`
	FromAcc AccountEntity `json:"FromAcc"`
}
