package twitter

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/config"
	"sync"
)

var once sync.Once

type twitterConfig struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

var instance *twitterConfig

func getConfigInstance() *twitterConfig {
	once.Do(func() {
		instance = &twitterConfig{
			consumerKey:    config.GetInstance().GetValue(ConsumerKey),
			consumerSecret: config.GetInstance().GetValue(ConsumerSecret),
			accessToken:    config.GetInstance().GetValue(AccessToken),
			accessSecret:   config.GetInstance().GetValue(AccessSecret),
		}
	})
	return instance
}
