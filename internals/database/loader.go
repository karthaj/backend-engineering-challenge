package database

import (
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/domain/log"
	"context"
	"fmt"
	"io/ioutil"
	"os"
)

const logPrefixLoader = "backend-engineering-challenge.internals.database.init"

func loadData(ctx context.Context) []byte {

	// Read JSON file into memory
	file, err := os.Open(config.AppConf.MockData)
	if err != nil {
		log.ErrorContext(ctx, logPrefixLoader, "Unable to load mock data")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.FatalContext(ctx, fmt.Sprintf("%s.%s", logPrefixLoader, "file.Close()"), err)
		}
	}(file)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.FatalContext(ctx, fmt.Sprintf("%s.%s", logPrefixLoader, "Unable Read mock data"), err)
	}

	return data
}
