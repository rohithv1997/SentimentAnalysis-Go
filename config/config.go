package applicationConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type configMap struct {
	keyValuePairs map[string]string
	rwMutex       sync.RWMutex
}

var Configuration configMap

func (configMap *configMap) GetValue(key string) string {
	configMap.rwMutex.RLock()
	defer configMap.rwMutex.RUnlock()
	return configMap.keyValuePairs[key]
}

func (configMap *configMap) initialize(key, value string) {
	configMap.rwMutex.Lock()
	defer configMap.rwMutex.Unlock()
	configMap.keyValuePairs[key] = value
}

var configJson = struct {
	Configs []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"configs"`
}{}

func LoadConfiguration() {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	if len(Configuration.keyValuePairs) != 0 {
		return
	}
	Configuration = configMap{
		keyValuePairs: make(map[string]string),
	}

	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "config", "config.json")

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &configJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, config := range configJson.Configs {
		Configuration.initialize(config.Key, config.Value)
	}
}
