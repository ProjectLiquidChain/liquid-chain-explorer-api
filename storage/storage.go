package storage

import (
	badger "github.com/dgraph-io/badger/v2"
)

const (
	TxStoragePath      = "txs"
	BlockStoragePath   = "blocks"
	ReceiptStoragePath = "receipts"
)

type Storage struct {
	db *badger.DB
}

func New(path string) Storage {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}
	return Storage{db}
}

func (storage *Storage) Update(key, value []byte) error {
	txn := storage.db.NewTransaction(true)
	if err := txn.Set(key, value); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) Get(key []byte) ([]byte, error) {
	var result []byte
	if err := storage.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		result, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}
