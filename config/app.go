package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Application struct {
	DB            *gorm.DB
	RabbitConn    *amqp.Connection
	RabbitChannel *amqp.Channel
}

func App() Application {
	app := &Application{}
	app.DB = NewGormConnection()
	app.RabbitConn, app.RabbitChannel = NewRabbitMQ()
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseGormConnection(app.DB)
}
