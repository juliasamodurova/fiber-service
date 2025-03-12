package config

import "time"

// общая конфигурация сервиса
const EnvPath = "local.env"

type AppConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
	Rest     Rest
}

type Rest struct {
	ListenAddress string        `envconfig:"PORT" required:"true"`
	WriteTimeout  time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	ServerName    string        `envconfig:"SERVER_NAME" required:"true"`
	Token         string        `envconfig:"TOKEN"`
}
