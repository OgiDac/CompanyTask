package config

import "gorm.io/gorm"

type Application struct {
	DB *gorm.DB
}

func App() Application {
	app := &Application{}
	app.DB = NewGormConnection()
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseGormConnection(app.DB)
}
