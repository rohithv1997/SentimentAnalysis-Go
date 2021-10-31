package config

var configJson = struct {
	Configs []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"configs"`
}{}
