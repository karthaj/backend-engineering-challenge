package entity

// Account struct represents the account information.
type Account struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

// AccountEntity struct represents the account information.
type AccountEntity struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
