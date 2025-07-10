package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	BaseDSN                string `mapstructure:"BASE_DSN"`
	TargetDB               string `mapstructure:"TARGET_DB"`
	RabbitMQUrl            string `mapstructure:"RABBITMQ_URL"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	MongoURL               string `mapstructure:"MONGO_URL"`
	MongoDBName            string `mapstructure:"MONGO_DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("CONTEXT_TIMEOUT")
	viper.BindEnv("BASE_DSN")
	viper.BindEnv("TARGET_DB")
	viper.BindEnv("RABBITMQ_URL")
	viper.BindEnv("ACCESS_TOKEN_EXPIRY_HOUR")
	viper.BindEnv("REFRESH_TOKEN_EXPIRY_HOUR")
	viper.BindEnv("ACCESS_TOKEN_SECRET")
	viper.BindEnv("REFRESH_TOKEN_SECRET")
	viper.BindEnv("MONGO_URL")
	viper.BindEnv("MONGO_DB_NAME")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No .env file found, relying on environment variables")
	}

	if err := viper.Unmarshal(&env); err != nil {
		fmt.Println("Environment can't be loaded:", err)
	}

	return &env
}
