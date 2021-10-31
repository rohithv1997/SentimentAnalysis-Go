package config

type data struct {
	keyValuePairs map[string]string
}

func (data *data) GetValue(key string) string {
	return data.keyValuePairs[key]
}

func (data *data) setKeyValuePair(key, value string) {
	data.keyValuePairs[key] = value
}
