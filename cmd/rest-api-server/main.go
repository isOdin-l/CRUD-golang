package main

import (
	"context"
	"net/http"

	"github.com/isOdin/RestApi/internal/database/postgresql"
	"github.com/isOdin/RestApi/internal/handler"
	"github.com/isOdin/RestApi/internal/httpchi"
	"github.com/isOdin/RestApi/internal/repository"
	"github.com/isOdin/RestApi/internal/service"
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
	r := httpchi.NewRouter(handler)

	// Server start
	server := httpchi.NewServer()
	go func() {
		if err := server.RunServer(viper.GetString("SERVER_PORT"), r); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error while running server %s", err.Error())
		}
	}()
	logrus.Print("Server started")

	server.GracefulShutdownServer(context.Background())
}
