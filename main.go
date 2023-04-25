package main

import (
	"backend-engineering-challenge/internals"
	"os"
)

func main() {

	// ******************************************
	// Dev test - interview running purpose only
	err := os.Setenv("DEBUG", "TRUE")
	err = os.Setenv("MOCK_DATA", "accounts-mock.json")
	err = os.Setenv("PORT", "8085")
	if err != nil {
		return
	}
	// ******************************************

	internals.Init()
}
