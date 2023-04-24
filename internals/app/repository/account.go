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

func (a AccountStruct) GetAllAccountDetails(_ context.Context) ([][]byte, error) {
	var val [][]byte

	err := database.Database.View(func(txn *badger.Txn) error {

		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			// Alternatively, you could also use item.ValueCopy().
			v, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			val = append(val, v)

		}

		return nil
	})

	return val, err
}

func (a AccountStruct) DoTransaction(_ context.Context, req entity.TransactionRequestEntity) (interface{}, error) {

	toData, err := json.Marshal(&req.ToAcc)
	if err != nil {
		return nil, err
	}

	fromData, err := json.Marshal(&req.FromAcc)
	if err != nil {
		return nil, err
	}

	// Insert the data into the database
	err = database.Database.Update(func(txn *badger.Txn) error {

		err = txn.Set([]byte(req.FromAcc.ID), fromData)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(req.ToAcc.ID), toData)
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
