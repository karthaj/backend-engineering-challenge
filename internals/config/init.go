package config

import (
	"os"
	"strconv"
)

var AppConf AppConfig

type AppConfig struct {
	Port     int
	Debug    bool
	MockData string
}

//	Load the config from environment
func InitConfig() {
	AppConf.Port = loadInt(os.Getenv("PORT"))
	AppConf.Debug = loadBool(os.Getenv("DEBUG"))
	AppConf.MockData = os.Getenv("MOCK_DATA")
}

// Load int value from string
func loadInt(v string) int {
	r, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		print("Error : ParseInt")
	}
	return int(r)
}

// loadFloat value from string
func loadFloat(v string) float64 {
	r, err := strconv.ParseFloat(v, 32)
	if err != nil {
		print("Error : ParseFloat")
	}
	return r
}

// Load bool value from string
func loadBool(v string) bool {
	r, err := strconv.ParseBool(v)
	if err != nil {
		print("Error : ParseBool")
	}
	return r
}
