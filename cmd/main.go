package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"log"
	"os"
	"os/signal"
	"sync"

	"fiber-service/internal/api"
	"fiber-service/internal/config"
	customLogger "fiber-service/internal/logger"
	"fiber-service/internal/repo"
	"fiber-service/internal/service"
)

func main() {
	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Error processing config", zap.Error(err))
	}

	logg, err := customLogger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal("Error initializing logger", zap.Error(err))
	}

	mu := new(sync.RWMutex)
	repos := repo.NewRepository(mu, logg)
	services := service.NewService(repos, logg)
	app := api.NewRouters(services.Task, logg)

	logg.Debug("Сервер запущен на порту :8080")
	if err := app.Listen(":8080"); err != nil {
		logg.Fatal("Ошибка запуска сервера", zap.Error(err))
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	logg.Info("Shutting down server...")
}
