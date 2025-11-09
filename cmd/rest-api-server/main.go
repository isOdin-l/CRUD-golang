package main

import (
	"context"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/isOdin/RestApi/configs"
	"github.com/isOdin/RestApi/internal/database/postgresql"
	"github.com/isOdin/RestApi/internal/handler"
	"github.com/isOdin/RestApi/internal/httpchi"
	"github.com/isOdin/RestApi/internal/middleware"
	"github.com/isOdin/RestApi/internal/repository"
	"github.com/isOdin/RestApi/internal/service"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	// Config
	var cfg configs.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Error("Error whie initialize config:, ", err.Error())
		return
	}
	internalCfg := &configs.InternalConfig{
		SALT:            cfg.SALT,
		JWT_SIGNING_KEY: cfg.JWT_SIGNING_KEY,
		TOKEN_TTL:       12 * time.Hour,
	}

	// Database: postgresql
	DB, err := postgresql.NewPostgresDB(&postgresql.Config{
		Host:     cfg.DB_HOST,
		Port:     cfg.DB_PORT,
		Username: cfg.DB_USERNAME,
		Password: cfg.DB_PASSWORD,
		DBName:   cfg.DB_NAME,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer DB.Close()

	// ----- Repository -----
	repository := repository.NewRepository(DB)

	// ----- Service -----
	service := service.NewService(internalCfg, repository)

	// ----- Validator -----
	validate := validator.New(validator.WithRequiredStructEnabled())

	// ----- Middleware -----
	middleware := middleware.NewMiddleware(internalCfg)

	// ----- Handler -----
	handler := handler.NewHandler(validate, service)

	// ----- Router -----
	r := httpchi.NewRouter(middleware, handler)

	// Server start
	server := httpchi.NewServer()
	go func() {
		if err := server.RunServer(cfg.SERVER_PORT, r); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error while running server %s", err.Error())
		}
	}()
	logrus.Print("Server started")

	server.GracefulShutdownServer(context.Background())
}
