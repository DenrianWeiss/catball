package service

import (
	"encoding/json"
	"errors"
	"github.com/DenrianWeiss/catball/model"
	"github.com/dgraph-io/badger/v3"
	"strings"
)

const (
	docsPrefix = "docs_"
)

var (
	DB        *badger.DB
	blackList = []string{"docs", "add", "show", "del"}
)

func AddRedirect(path, target string) error {
	for _, r := range blackList {
		if strings.HasPrefix(path, r) {
			return errors.New("banned path")
		}
	}
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

func DelRedirect(path string) error {
	return DB.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(path))
		if err == badger.ErrKeyNotFound {
			return errors.New("no path to delete")
		} else {
			return txn.Delete([]byte(path))
		}
	})
}

func AddDocument(path string, article *model.Article) error {
	return DB.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(docsPrefix + path))
		if err == badger.ErrKeyNotFound {
			j, err := json.Marshal(article)
			if err != nil {
				return err
			}
			return txn.Set([]byte(docsPrefix+path), j)
		}
		return err
	})
}

func GetDocument(path string, result *model.Article) error {
	return DB.View(func(txn *badger.Txn) error {
		res, err := txn.Get([]byte(docsPrefix + path))
		if err != nil {
			return err
		}
		bytes, err := res.ValueCopy(nil)
		if err != nil {
			return err
		}
		article := &model.Article{}
		err = json.Unmarshal(bytes, article)
		if err != nil {
			return err
		}
		*result = *article
		return nil
	})
}

func DelDocument(path string) error {
	return DB.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(docsPrefix + path))
	})
}
