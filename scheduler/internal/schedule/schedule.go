package schedule

import (
	"calendar/entities"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Serve()  {
	conn, err := amqp.Dial(viper.GetString("ampq"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer amqpChannel.Close()

	q, err := amqpChannel.QueueDeclare(
		"events", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	stopChan := make(chan bool)

	ticker := time.NewTicker(time.Second * 5)

	go func() {
		for t := range ticker.C {
			fmt.Println("Check event", t)
			checkEvent(amqpChannel, q)
		}
	}()

	<-stopChan

	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, os.Interrupt)
}

func checkEvent(amqpChannel *amqp.Channel, q amqp.Queue)  {


	requestCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Establish insecure grpc options (no TLS)
	requestOpts := grpc.WithInsecure()

	conn, err := grpc.DialContext(requestCtx, ":50051", requestOpts)

	if err != nil {
		log.Fatalf("Unable to establish client connection to : %s, %v", ":50051", err)
	}

	client := entities.NewEventServiceClient(conn)

	req := &entities.EventListRequest{
		Period: entities.Period_DAY,
	}

	stream, err := client.EventList(context.Background(), req)

	if err != nil {
		failOnError(err, "Failed to read steam")
	}

	for {
		// stream.Recv returns a pointer to a EventList at the current iteration
		res, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		// if err, return an error
		if err != nil {
			failOnError(err, "Failed to read steam")
		}

		body, err := json.Marshal(res.Event)

		err = amqpChannel.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")

		// If everything went well use the generated getter to print the event message
		fmt.Println(res.Event)
	}

}