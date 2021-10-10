package applicationConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type configurationManager struct {
	keyValuePairs map[string]string
}

var once sync.Once

var instance *configurationManager

func (configMap *configurationManager) GetValue(key string) string {
	return configMap.keyValuePairs[key]
}

func (configMap *configurationManager) setKeyValuePair(key, value string) {
	configMap.keyValuePairs[key] = value
}

var configJson = struct {
	Configs []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"configs"`
}{}

func GetInstance() *configurationManager {
	once.Do(func() {
		instance = &configurationManager{
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
