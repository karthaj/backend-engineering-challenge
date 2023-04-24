package database

import (
	"backend-engineering-challenge/internals/domain/entity"
	"backend-engineering-challenge/internals/domain/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	log2 "log"
	"os"
	"time"
)

const logPrefixDatabase = "backend-engineering-challenge.internals.database.init"

type BadgerDB struct {
	db         *badger.DB
	ctx        context.Context
	cancelFunc context.CancelFunc
	logger     *log2.Logger
}

var DbCon *BadgerDB

var Database *badger.DB

const (
	// Default BadgerDB discardRatio. It represents the discard ratio for the
	// BadgerDB GC.
	//
	// Ref: https://godoc.org/github.com/dgraph-io/badger#DB.RunValueLogGC
	badgerDiscardRatio = 0.5

	// Default BadgerDB GC interval
	badgerGCInterval = 10 * time.Minute
)

func Init(ctx context.Context) {

	dirPath := "/tmp/badger"
	if err := os.MkdirAll("/tmp/badger", 0774); err != nil {
		log.ErrorContext(ctx, logPrefixDatabase, err)
	}

	opts := badger.DefaultOptions("/tmp/badger")
	opts.SyncWrites = true
	opts.Dir, opts.ValueDir = dirPath, dirPath

	badgerDB, err := badger.Open(opts)
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "badger.Open(opts)"), err)
	}

	Database = badgerDB

	bdb := &BadgerDB{
		db:     badgerDB,
		logger: log.NativeLog,
		ctx:    ctx,
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(ctx)

	go bdb.runGC(ctx)

	DbCon = bdb

}

func LoadDB(ctx context.Context) error {

	err := Database.DropAll()
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "r.db.DropAll"), "Unable to unmarshal data", fmt.Sprintf("%+v", err.Error()))
		return err
	}

	// Unmarshal JSON data into Account slice
	var accounts []entity.Account
	err = json.Unmarshal(loadData(ctx), &accounts)
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "json.Unmarshal(loadData(ctx), &accounts)"), "Unable to unmarshal data", fmt.Sprintf("%+v", err.Error()))
	}

	txn := Database.NewWriteBatch()
	defer txn.Cancel()

	// Add each key-value pair to the batch
	for _, d := range accounts {
		key := []byte(d.ID)

		data, err := json.Marshal(d)
		if err != nil {
			log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "json.Marshal(v)"), "failed to parse data")
			return err
		}
		value := data
		err = txn.Set(key, value)
		if err != nil {
			log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "json.Marshal(v)"), "failed to insert data")

		}
	}

	// Commit the batch write transaction
	err = txn.Flush()
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "json.Marshal(v)"), "failed to load data in to the DB")
	}

	log.InfoContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "LoadDB"), "Mock data loaded in to the DB successfully")

	return nil
}

// Get implements the DB interface. It attempts to get a value for a given key
// and namespace. If the key does not exist in the provided namespace, an error
// is returned, otherwise the retrieved value.
func (r *BadgerDB) Get(key string) (value []byte, err error) {

	var val []byte
	err = r.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(key))
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

//
//// Set implements the DB interface. It attempts to store a value for a given key
//// and namespace. If the key/value pair cannot be saved, an error is returned.
//func  Set(namespace, key string, v interface{}) error {
//
//	// Encode the data as JSON
//	data, err := json.Marshal(v)
//	if err != nil {
//		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "json.Marshal(v)"), "failed to set key %s for namespace %s: %v", key, namespace, err)
//		return err
//	}
//
//	// Insert the data into the database
//	err = r.db.Update(func(txn *badger.Txn) error {
//		err := txn.Set([]byte(key), data)
//		return err
//	})
//
//	if err != nil {
//		log.DebugContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "r.db.Update()"), "failed to set key %s for namespace %s: %v", key, namespace, err)
//		return err
//	}
//
//	return nil
//}

// Has implements the DB interface. It returns a boolean reflecting if the
// database has a given key for a namespace or not. An error is only returned if
// an error to Get would be returned that is not of type badger.ErrKeyNotFound.
func (r *BadgerDB) Has(key string) (ok bool, err error) {
	_, err = r.Get(key)
	switch err {
	case badger.ErrKeyNotFound:
		ok, err = false, nil
	case nil:
		ok, err = true, nil
	}

	return false, nil
}

//// Close implements the DB interface. It closes the connection to the underlying
//// BadgerDB database as well as invoking the context's cancel function.
//func (r *BadgerDB) Close() error {
//	r.cancelFunc()
//	return r.db.Close()
//}

// badgerNamespaceKey returns a composite key used for lookup and storage for a
// given namespace and key.
func badgerNamespaceKey(namespace, key []byte) []byte {
	prefix := []byte(fmt.Sprintf("%s/", namespace))
	return append(prefix, key...)
}

// runGC triggers the garbage collection for the BadgerDB backend database. It
// should be run in a goroutine.
func (r *BadgerDB) runGC(ctx context.Context) {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			{
				err := r.db.RunValueLogGC(badgerDiscardRatio)
				if err != nil {
					// don't report error when GC didn't result in any cleanup

					if err == badger.ErrNoRewrite {
						log.DebugContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "jr.db.RunValueLogGC(badgerDiscardRatio)"), "no BadgerDB GC occurred", err.Error())
					} else {
						log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "jr.db.RunValueLogGC(badgerDiscardRatio)"), "failed to GC BadgerDB", err.Error())
					}
				}
				log.InfoContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "Cleaning"), "no BadgerDB GC occurred:")

			}

		case <-ctx.Done():
			log.InfoContext(ctx, fmt.Sprintf("%s.%s", logPrefixDatabase, "Cleaning"), "no BadgerDB GC occurred")

			return
		}
	}
}
