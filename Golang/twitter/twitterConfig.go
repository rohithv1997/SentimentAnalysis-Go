package twitter

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/applicationConfig"
	"sync"
)

const (
	prefix         = "Twitter_"
	ConsumerKey    = prefix + "ConsumerKey"
	ConsumerSecret = prefix + "ConsumerSecret"
	AccessToken    = prefix + "AccessToken"
	AccessSecret   = prefix + "AccessSecret"
)

var once sync.Once

type twitterConfig struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

var instance *twitterConfig

func GetTwitterConfigInstance() *twitterConfig {
	once.Do(func() {
		instance = &twitterConfig{
			consumerKey:    applicationConfig.GetInstance().GetValue(ConsumerKey),
			consumerSecret: applicationConfig.GetInstance().GetValue(ConsumerSecret),
			accessToken:    applicationConfig.GetInstance().GetValue(AccessToken),
			accessSecret:   applicationConfig.GetInstance().GetValue(AccessSecret),
		}
	})
	return instance
}
