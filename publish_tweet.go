package main

import (
	"github.com/google/uuid"
	"github.com/rohithv1997/SentimentAnalysis-Go/models"
	"github.com/rohithv1997/SentimentAnalysis-Go/rabbitmq"
)

func PublishTweet(message string) {
	payload := models.OutgoingMessage{
		Message:   message,
		MessageId: uuid.New().String(),
	}
	rabbitmq.Publish(payload)
}
