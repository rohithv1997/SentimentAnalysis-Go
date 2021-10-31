package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once

var instance *data

func GetInstance() *data {
	once.Do(func() {
		instance = &data{
			keyValuePairs: make(map[string]string),
		}

		pwd, _ := os.Getwd()
		path := filepath.Join(pwd, "config", "config.json")

		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(jsonFile)

		byteValue, _ := ioutil.ReadAll(jsonFile)

		err = json.Unmarshal(byteValue, &configJson)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, config := range configJson.Configs {
			instance.setKeyValuePair(config.Key, config.Value)
		}
	})

	return instance
}
