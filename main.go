package main

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/applicationConfig"
	"github.com/rohithv1997/SentimentAnalysis-Go/twitter"
)

func main() {
	applicationConfig.LoadConfiguration()
	twitter.ApiEndpoint("olympics")
}
