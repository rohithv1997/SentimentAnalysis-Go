package main

import (
	"github.com/google/uuid"
	"github.com/rohithv1997/SentimentAnalysis-Go/models"
	"github.com/rohithv1997/SentimentAnalysis-Go/rabbitmq"
	"github.com/rohithv1997/SentimentAnalysis-Go/twitter"
)

func main() {
	twitter.StreamApi("india", publishTweet)
}

func publishTweet(message string) {
	payload := models.OutgoingMessage{
		Message:   message,
		MessageId: uuid.NewString(),
	}
	rabbitmq.Publish(payload)
}
