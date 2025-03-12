package config

// общая конфигурация сервиса

type AppConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL"`
}
