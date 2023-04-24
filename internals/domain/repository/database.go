package repository

import (
	"context"
	"github.com/dgraph-io/badger/v2"
	log2 "log"
)

type DB interface {
	Get(namespace, key []byte) (value []byte, err error)
	Set(namespace, key, value []byte) error
	Has(namespace, key []byte) (bool, error)
	Close() error
}

type BadgerDB struct {
	db         *badger.DB
	ctx        context.Context
	cancelFunc context.CancelFunc
	logger     *log2.Logger
}
