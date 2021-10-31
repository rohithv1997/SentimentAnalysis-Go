package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/rohithv1997/SentimentAnalysis-Go/models"
	"github.com/streadway/amqp"
	"log"
)

func Publish(payload models.OutgoingMessage) {
	message, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	url := fmt.Sprintf(urlTemplate,
		GetRabbitMqConfigInstance().username,
		GetRabbitMqConfigInstance().password,
		GetRabbitMqConfigInstance().url)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Println(err)
		}
	}(ch)

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		GetRabbitMqConfigInstance().publishQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		log.Println(err)
	}

	err = ch.ExchangeDeclare(
		GetRabbitMqConfigInstance().exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err)
		return
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		GetRabbitMqConfigInstance().exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        message,
		},
	)

	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Published Message to Queue")
}
