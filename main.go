package main

import "github.com/rohithv1997/SentimentAnalysis-Go/twitter"

func main() {
	twitter.StreamApi("india", func(message string) {
		PublishTweet(message)
	})
}
