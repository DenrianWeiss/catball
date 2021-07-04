package service

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
)

var (
	DB *badger.DB
)

func AddRedirect(path, target string) error {
	return DB.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(path))
		if err == badger.ErrKeyNotFound {
			return txn.Set([]byte(path), []byte(target))
		}
		return errors.New("exists")
	})
}

func GetRedirect(path string, result *string) error {
	return DB.View(func(txn *badger.Txn) error {
		res, err := txn.Get([]byte(path))
		if err != nil {
			return err
		}
		bytes, err := res.ValueCopy(nil)
		if err != nil {
			return err
		}
		b := string(bytes)
		*result = b
		return nil
	})
}

