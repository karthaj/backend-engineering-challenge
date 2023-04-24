package repository

import (
	"backend-engineering-challenge/internals/database"
	"backend-engineering-challenge/internals/domain/entity"
	"context"
	"encoding/json"
	"github.com/dgraph-io/badger/v2"
)

type AccountStruct struct{}

var AccountRepository = AccountStruct{}

func (a AccountStruct) GetAccountDetailsByID(ctx context.Context, id string) ([]byte, error) {

	var val []byte

	err := database.Database.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		// Alternatively, you could also use item.ValueCopy().
		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	return val, err
}

func (a AccountStruct) GetAccountDetailsByName(_ context.Context, _ string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountStruct) DoTransaction(_ context.Context, req entity.AccountEntity) (interface{}, error) {

	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	// Insert the data into the database
	err = database.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set([]byte(req.ID), data)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil

}
