package repository

import (
	"backend-engineering-challenge/internals/database"
	"backend-engineering-challenge/internals/domain/entity"
	req_res "backend-engineering-challenge/internals/domain/req-res"
	"context"
	"github.com/dgraph-io/badger/v2"
)

type AccountStruct struct{}

var AccountRepository = AccountStruct{}

func (a AccountStruct) GetAccountDetailsByID(_ context.Context, id string) ([]byte, error) {

	var val []byte
	err := database.Database.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(b []byte) error {
			// This func with val would only be called if item.Value encounters no error.
			// Copying or parsing val is valid.
			val = append([]byte{}, b...)

			return nil
		})

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

func (a AccountStruct) GetAccountDetailsByName(_ context.Context, name string) ([]byte, error) {

	var val []byte
	err := database.Database.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(name))
		if err != nil {
			return err
		}

		err = item.Value(func(b []byte) error {
			// This func with val would only be called if item.Value encounters no error.
			// Copying or parsing val is valid.
			val = append([]byte{}, b...)

			return nil
		})

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

func (a AccountStruct) DoTransaction(ctx context.Context, data req_res.DoTransactionRequest) (entity.Account, error) {
	//TODO implement me
	panic("implement me")
}
