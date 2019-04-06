package infrastructure

import (
	"errors"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMqHandler struct {
	conn *amqp.Connection
}

func (handler *RabbitMqHandler) PublishOnQueue(msg []byte, queueName string) error {
	if handler.conn == nil {
		return errors.New("Tried to send message before connection was initialized")
	}

	channel, err := handler.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Message %v successfully published on queue %v\n", string(msg), queue.Name)
	return nil
}

func NewRabbitMqHandler(username, password, hostname, port string) (*RabbitMqHandler, error) {
	connString := fmt.Sprintf("amqp://%v:%v@%v:%v", username, password, hostname, port)
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stdout, "Connected to RabbitMQ server\n")

	rabbitMqHandler := new(RabbitMqHandler)
	rabbitMqHandler.conn = conn
	return rabbitMqHandler, nil
}
