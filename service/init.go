package service

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v3"
	"io/ioutil"
	"os"
)

func init() {
	configFile, err := os.Open(config)
	if err != nil {
		panic("No config.")
	}
	configContent, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic("Failed to read config")
	}

	configResult := GlobalConfig{}

	err = json.Unmarshal(configContent, &configResult)

	if err != nil {
		panic("Illegal config")
	}

	Config = configResult

	db, err := badger.Open(badger.DefaultOptions(Config.DatabaseFile))
	if err != nil {
		panic("Cannot open database")
	}
	DB = db
}
