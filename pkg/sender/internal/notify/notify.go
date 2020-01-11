package notify

import (
	"encoding/json"
	"github.com/assizkii/calendar/entities"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Serve() {
	conn, err := amqp.Dial(viper.GetString("ampq"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("events", false, false, false, false, nil)
	failOnError(err, "Could not declare `events` queue")

	sendNotify(amqpChannel, queue)
}

func sendNotify(amqpChannel *amqp.Channel, queue amqp.Queue) {
	err := amqpChannel.Qos(1, 0, false)
	failOnError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		for d := range messageChannel {
			event := &entities.Event{}
			err := json.Unmarshal(d.Body, event)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("You have event today: %s", event.Title)
			}

		}
	}()
	// Stop for program termination
	<-stopChan
}
