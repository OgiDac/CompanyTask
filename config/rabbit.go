package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		"user-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	return conn, ch
}

func (app *Application) CloseRabbitConnection() {
	if app.RabbitChannel != nil {
		_ = app.RabbitChannel.Close()
	}
	if app.RabbitConn != nil {
		_ = app.RabbitConn.Close()
	}
}
