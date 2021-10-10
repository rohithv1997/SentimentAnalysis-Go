package rabbitmq

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/applicationConfig"
	"sync"
)

const (
	prefix         = "RabbitMq_"
	Url            = prefix + "Url"
	Username       = prefix + "Username"
	Password       = prefix + "Password"
	Exchange       = prefix + "Exchange"
	PublishQueue   = prefix + "PublishQueue"
	SubscribeQueue = prefix + "SubscribeQueue"
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

func GetRabbitMqConfigInstance() *rabbitMqConfig {
	once.Do(func() {
		instance = &rabbitMqConfig{
			url:            applicationConfig.GetInstance().GetValue(Url),
			username:       applicationConfig.GetInstance().GetValue(Username),
			password:       applicationConfig.GetInstance().GetValue(Password),
			exchange:       applicationConfig.GetInstance().GetValue(Exchange),
			publishQueue:   applicationConfig.GetInstance().GetValue(PublishQueue),
			subscribeQueue: applicationConfig.GetInstance().GetValue(SubscribeQueue),
		}
	})
	return instance
}
