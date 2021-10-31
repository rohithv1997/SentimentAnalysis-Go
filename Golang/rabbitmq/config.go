package rabbitmq

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/config"
	"sync"
)

var once sync.Once

type rabbitMqConfig struct {
	url            string
	username       string
	password       string
	exchange       string
	publishQueue   string
	subscribeQueue string
}

var instance *rabbitMqConfig

func getConfigInstance() *rabbitMqConfig {
	once.Do(func() {
		instance = &rabbitMqConfig{
			url:            config.GetInstance().GetValue(Url),
			username:       config.GetInstance().GetValue(Username),
			password:       config.GetInstance().GetValue(Password),
			exchange:       config.GetInstance().GetValue(Exchange),
			publishQueue:   config.GetInstance().GetValue(PublishQueue),
			subscribeQueue: config.GetInstance().GetValue(SubscribeQueue),
		}
	})
	return instance
}
