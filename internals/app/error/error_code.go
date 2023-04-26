package errors

const (
	//ErrTokenInvalid = 100001
	ErrAuthNoToken = 100002 // No Token | Authenticator
)

const (
	// ErrResponseParse is an error code for failed response parsing
	ErrResponseParse = 100102

	// ErrParamNotFound is an error code for missing parameter
	ErrParamNotFound = 100103

	// ErrBatabaseRead is an error code for failed database read operation
	ErrBatabaseRead = 100104

	// ErrAccountNotFound is an error code for account not found
	ErrAccountNotFound = 100105

	// ErrAccountBalanceInsufficient is an error code for insufficient account balance
	ErrAccountBalanceInsufficient = 100106

	// ErrTransactionFailed is an error code for failed transaction
	ErrTransactionFailed = 100106

	// ErrNegativeTrxAmount is an error code for negative transfer amount
	ErrNegativeTrxAmount = 100107

	// ErrRequestInvalid is an error code for invalid request
	ErrRequestInvalid = 100101
)
