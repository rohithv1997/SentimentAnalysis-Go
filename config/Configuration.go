package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Configuration struct {
	key   string
	value string
}

type configurations struct {
	configs []Configuration
}

var configs configurations

func LoadConfiguration() {
	if reflect.DeepEqual(configs, configurations{}) {
		jsonFile, err := os.Open("config.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &configs)
	}

}
