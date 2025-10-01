package main

import (
	"fmt"
	"log"

	"github.com/isOdin/RestApi/api/handler"
	"github.com/isOdin/RestApi/internal/router"
	"github.com/isOdin/RestApi/internal/server"
	"github.com/isOdin/RestApi/internal/storage/postgresql"
	"github.com/isOdin/RestApi/pkg/repository"
	"github.com/isOdin/RestApi/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// Database: postgresql
	DB, err := postgresql.NewPostgresDB(&postgresql.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer DB.Close()

	// Repositories
	repository := repository.NewRepository(DB)

	// Services
	service := service.NewService(repository)

	// Handlers
	authHandler := handler.NewAuthHandler(service)
	itemHandler := handler.NewItemHandler(service)
	listHandler := handler.NewListHandler(service)

	// Router
	r := router.NewRouter(listHandler, itemHandler, authHandler)

	// Server start
	server := server.New()
	if err := server.Run(viper.GetString("server.port"), r); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}

func initConfig() error {
	if err := godotenv.Load("./configs/.env"); err != nil {
		fmt.Println(".env file not found, using environment variables only")
		return err
	}

	viper.AutomaticEnv()

	return nil
}
