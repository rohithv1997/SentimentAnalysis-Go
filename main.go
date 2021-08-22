package main

import (
	"github.com/rohithv1997/SentimentAnalysis-Go/applicationConfig"
	"github.com/rohithv1997/SentimentAnalysis-Go/corenlp"
)

func main() {
	applicationConfig.LoadConfiguration()
	//twitter.ApiEndpoint("olympics")
	corenlp.Process()
}
