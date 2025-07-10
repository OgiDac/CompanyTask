package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Application struct {
	DB            *gorm.DB
	RabbitConn    *amqp.Connection
	RabbitChannel *amqp.Channel
	Env           *Env
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewGormConnection(app.Env)
	app.RabbitConn, app.RabbitChannel = NewRabbitMQ(app.Env)
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseGormConnection(app.DB)
}
