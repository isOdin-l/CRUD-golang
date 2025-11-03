package main

import (
	"context"

	"github.com/isOdin/RestApi/api/handler"
	"github.com/isOdin/RestApi/internal/router"
	"github.com/isOdin/RestApi/internal/server"
	"github.com/isOdin/RestApi/internal/storage/postgresql"
	"github.com/isOdin/RestApi/pkg/repository"
	"github.com/isOdin/RestApi/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	viper.AutomaticEnv()
}

func main() {
	// Database: postgresql
	DB, err := postgresql.NewPostgresDB(&postgresql.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer DB.Close()

	// Repository
	repository := repository.NewRepository(DB)

	// Service
	service := service.NewService(repository)

	// Handler
	handler := handler.NewHandler(service)

	// Router
	r := router.NewRouter(handler)

	// Server start
	server := server.New()
	go func() {
		if err := server.Run(viper.GetString("SERVER_PORT"), r); err != nil {
			logrus.Fatalf("error while running server %s", err.Error())
		}
	}()
	logrus.Print("Server started")

	server.GracefulShutdown(context.Background())
}
